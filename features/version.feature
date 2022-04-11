Feature: get version
  In order to know service version
  As an API user
  I need to be able to request version

  Scenario Outline: should get version number
    When I send GET request to /version
    Then the response code should be <status>
    And the response should match json:
      """
      {
        "version": "v2.12.3"
      }
      """
    Examples:
      | status |
      | 200    |