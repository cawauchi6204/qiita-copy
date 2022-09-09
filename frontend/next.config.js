/** @type {import('next').NextConfig} */
const withInterceptStdout = require('next-intercept-stdout');

const warningIgnoreConfig = withInterceptStdout(
  {
    reactStrictMode: true,
    swcMinify: false,
    ignoreDuringBuilds: true,
    staticPageGenerationTimeout: 1000
  },
  (text) => (text.includes('Duplicate atom key') ? '' : text),
);

module.exports = {
  reactStrictMode: true,
  ...warningIgnoreConfig
}
