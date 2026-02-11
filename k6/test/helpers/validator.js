import { expect } from 'https://jslib.k6.io/k6chaijs/4.3.4.3/index.js';

export function createValidator(scopeName) {
  return (actualValue, checkName) => {
    const name = `[${scopeName}] ${checkName}`;
    return expect(actualValue, name);
  };
}