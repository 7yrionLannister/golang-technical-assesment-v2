import encoding from 'k6/encoding';

const SLEEP_DURATION = 0.1;

const BASE_URL = __ENV.BASE_URL || 'http://localhost:8181';
const PORT = __ENV.TENANTS || '8181';
const DB_HOST = __ENV.TENANTS || 'localhost:5432';
const DB_USER = __ENV.TENANTS || 'dev-user';
const DB_PASSWORD = __ENV.TENANTS || 'dev-password';
const DB_NAME = __ENV.TENANTS || 'meters';
const DB_ENGINE = __ENV.TENANTS || 'postgresql';
const LOG_LEVEL = __ENV.TENANTS || 'debug';
const TOKEN = __ENV.TOKEN || 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.KMUFsIDTnFmyG3nMiGM6H9FNFUROf3wh7SmqJp-QV30'

export { SLEEP_DURATION, BASE_URL, PORT, DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_ENGINE, LOG_LEVEL, TOKEN };