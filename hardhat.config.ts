import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";

const config: HardhatUserConfig = {
  solidity: "0.8.19",
  mocha: {
    reporter: "json",
    reporterOptions: {
      output: "test-results.json",
    },
  },
  defaultNetwork: "localhost",
};

export default config;
