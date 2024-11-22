// scripts/fetch-and-convert-schema.js
const fs = require('fs');
const path = require('path');
const http = require('http');
const https = require('https');
const { buildClientSchema, getIntrospectionQuery, printSchema } = require('graphql');
const { URL } = require('url');

const SCHEMA_DIR = 'internal/adapters/secondary/piyographql/schema';
const SCHEMA_JSON_PATH = path.join(SCHEMA_DIR, 'schema.json');
const SCHEMA_GRAPHQL_PATH = path.join(SCHEMA_DIR, 'schema.graphql');

// フルイントロスペクションクエリを使用
const INTROSPECTION_QUERY = getIntrospectionQuery();

async function fetchSchema() {
    const endpoint = process.env.GRAPHQL_ENDPOINT;
    if (!endpoint) {
        throw new Error('GRAPHQL_ENDPOINT environment variable is not set');
    }

    console.log('📡 Fetching schema from:', endpoint);

    return new Promise((resolve, reject) => {
        const url = new URL(endpoint);
        const requestOptions = {
            method: 'POST',
            hostname: url.hostname,
            port: url.port || (url.protocol === 'https:' ? 443 : 80),
            path: url.pathname + url.search,
            headers: {
                'Content-Type': 'application/json',
            },
        };

        // プロトコルに応じてクライアントを選択
        const client = url.protocol === 'https:' ? https : http;

        const req = client.request(requestOptions, (res) => {
            let data = '';

            res.on('data', (chunk) => {
                data += chunk;
            });

            res.on('end', () => {
                if (res.statusCode !== 200) {
                    reject(new Error(`HTTP Error: ${res.statusCode} ${res.statusMessage}`));
                    return;
                }

                try {
                    const jsonData = JSON.parse(data);
                    resolve(jsonData);
                } catch (error) {
                    reject(new Error(`Failed to parse JSON response: ${error.message}`));
                }
            });
        });

        req.on('error', (error) => {
            reject(new Error(`Request failed: ${error.message}`));
        });

        const postData = JSON.stringify({ query: INTROSPECTION_QUERY });
        req.write(postData);
        req.end();
    });
}

async function convertSchema(schemaJson) {
    try {
        console.log('🔄 Converting schema to GraphQL SDL...');
        const schema = buildClientSchema(schemaJson.data);
        return printSchema(schema);
    } catch (error) {
        throw new Error(`Schema conversion failed: ${error.message}`);
    }
}

async function ensureDirectoryExists(dirPath) {
    if (!fs.existsSync(dirPath)) {
        fs.mkdirSync(dirPath, { recursive: true });
    }
}

async function main() {
    try {
        // ディレクトリの存在確認
        await ensureDirectoryExists(SCHEMA_DIR);

        // スキーマのフェッチ
        console.log('🚀 Starting schema fetching and conversion process...');
        const schemaJson = await fetchSchema();

        // JSONスキーマの保存
        console.log('💾 Saving JSON schema...');
        fs.writeFileSync(SCHEMA_JSON_PATH, JSON.stringify(schemaJson, null, 2));

        // GraphQLスキーマへの変換と保存
        const graphqlSchema = await convertSchema(schemaJson);
        console.log('💾 Saving GraphQL schema...');
        fs.writeFileSync(SCHEMA_GRAPHQL_PATH, graphqlSchema);

        console.log('✅ Schema fetch and conversion completed successfully!');
        console.log(`📁 JSON schema saved to: ${SCHEMA_JSON_PATH}`);
        console.log(`📁 GraphQL schema saved to: ${SCHEMA_GRAPHQL_PATH}`);
    } catch (error) {
        console.error('❌ Error:', error.message);
        process.exit(1);
    }
}

main();
