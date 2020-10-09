pragma solidity ^0.7.1;

contract Store {

  string private scoreJson = "{}";

  function getScoreJson() public view returns (string memory)
  {
    return scoreJson;
  }

  function setScoreJson(string calldata newScoreJson) public
  {
    scoreJson = newScoreJson;
  }
}
