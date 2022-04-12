@version
Feature: Version
  In order to know service version
  As an API user
  I need to be able to request version

  Scenario Outline: should get version number
    When client send "GET" request to "/version"
    Then the response code should be <status>
    And the response should match json:
      """
      {
        "version": "<version>"
      }
      """
    Examples:
      | status | version |
      | 200    | v1.0.0  |