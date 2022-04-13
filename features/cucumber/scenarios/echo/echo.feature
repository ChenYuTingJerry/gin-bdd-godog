@api @echo
Feature: Echo API
  Echo API

  Scenario Outline: Should response what client send
    Given client give a request body:
    """
      {
        "echo": "hello world"
      }
    """
    When client send "POST" request to "/echo"
    Then the response code should be <http_status>
    And the response should match json:
    """
      {
        "echo": "hello world"
      }
    """
    Examples:
      | http_status |
      | 200         |
