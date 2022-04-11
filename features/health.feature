Feature: check health
  In order to know service health
  As an API maintainer
  I need to know the service health status

  Scenario Outline: should get ok status
    When I send GET request to /health
    Then the response code should be <status>
    And the response should match json:
      """
      {
        "status": "ok"
      }
      """
    Examples:
      | status |
      | 200    |