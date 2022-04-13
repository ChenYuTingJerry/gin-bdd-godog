@api @health
Feature: Health API
  In order to know service health
  As an API maintainer
  I need to know the service health status

  Scenario Outline: should get ok status
    Given
    When client send "GET" request to "/health"
    Then the response code should be <http_status>
    And the response should match json:
      """
      {
        "status": "<health_status>"
      }
      """
    Examples:
      | http_status | health_status |
      | 200         | ok            |
