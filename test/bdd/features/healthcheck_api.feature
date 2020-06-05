#
# Copyright SecureKey Technologies Inc. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

@all
@healthcheck
Feature: health check

  Scenario Outline:
    When an HTTP GET is sent to "<url>"
    Then the JSON path "<respKey>" of the response equals "<respKeyVal>"
    Examples:
      | url                                             | respKey       | respKeyVal                                      |
      | http://localhost:8070/healthcheck               | status        | success                                         |
      | http://localhost:8060/healthcheck               | status        | success                                         |
      | http://localhost:8060/issuer/didcomm/invitation | @type         | https://didcomm.org/didexchange/1.0/invitation  |
