module.exports = {
  testEnvironment: "jsdom",
  transform: {
    "^.+\\.(css|scss|sass)$": "jest-preview/transforms/css",
    "^(?!.*\\.(js|jsx|mjs|cjs|ts|tsx|css|json)$)":
      "jest-preview/transforms/file",
  },
  setupFilesAfterEnv: ['<rootDir>/src/setupTests.ts'],
};
