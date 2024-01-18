Feature: Test mTLS between client and gateway with client certificate sent in header
    Scenario: Test mandatory mTLS and enabled OAuth2 with a valid client certificate in header
        Given The system is ready
        And I have a valid token with a client certificate "config-map-1.txt"
        When I use the APK Conf file "artifacts/apk-confs/mtls/mtls_mandatory_oauth2_enabled.apk-conf"
        And the definition file "artifacts/definitions/employees_api.json"
        And make the API deployment request
        Then the response status code should be 200
        Then I set headers
            | Authorization             | bearer ${accessToken} |
            | X-WSO2-CLIENT-CERTIFICATE | ${clientCertificate}  |
        And I send "GET" request to "https://default.gw.wso2.com:9095/mtls/3.14/employee/" with body ""
        And I eventually receive 200 response code, not accepting
            | 401 |
        When I undeploy the API whose ID is "mtls-mandatory-oauth2-enabled"
        Then the response status code should be 202

    Scenario: Test mandatory mTLS and enabled OAuth2 with an invalid client certificate in header
        Given The system is ready
        And I have a valid token with a client certificate "invalid-cert.txt"
        When I use the APK Conf file "artifacts/apk-confs/mtls/mtls_mandatory_oauth2_enabled.apk-conf"
        And the definition file "artifacts/definitions/employees_api.json"
        And make the API deployment request
        Then the response status code should be 200
        Then I set headers
            | Authorization             | bearer ${accessToken} |
            | X-WSO2-CLIENT-CERTIFICATE | ${clientCertificate}  |
        And I send "GET" request to "https://default.gw.wso2.com:9095/mtls/3.14/employee/" with body ""
        And I eventually receive 401 response code, not accepting
            | 200 |
        When I undeploy the API whose ID is "mtls-mandatory-oauth2-enabled"
        Then the response status code should be 202

    Scenario: Test mandatory mTLS and enabled OAuth2 without a client certificate in header
        Given The system is ready
        And I have a valid subscription
        When I use the APK Conf file "artifacts/apk-confs/mtls/mtls_mandatory_oauth2_enabled.apk-conf"
        And the definition file "artifacts/definitions/employees_api.json"
        And make the API deployment request
        Then the response status code should be 200
        Then I set headers
            | Authorization | bearer ${accessToken} |
        And I send "GET" request to "https://default.gw.wso2.com:9095/mtls/3.14/employee/" with body ""
        And I eventually receive 401 response code, not accepting
            | 200 |
        When I undeploy the API whose ID is "mtls-mandatory-oauth2-enabled"
        Then the response status code should be 202

    Scenario: Test mandatory mTLS and disabled OAuth2 with a valid client certificate in header
        Given The system is ready
        And I have a valid token with a client certificate "config-map-1.txt"
        When I use the APK Conf file "artifacts/apk-confs/mtls/mtls_mandatory_oauth2_disabled.apk-conf"
        And the definition file "artifacts/definitions/employees_api.json"
        And make the API deployment request
        Then the response status code should be 200
        Then I set headers
            | X-WSO2-CLIENT-CERTIFICATE | ${clientCertificate} |
        And I send "GET" request to "https://default.gw.wso2.com:9095/mtls/3.14/employee/" with body ""
        And I eventually receive 200 response code, not accepting
            | 401 |
        When I undeploy the API whose ID is "mtls-mandatory-oauth2-disabled"
        Then the response status code should be 202

    Scenario: Test mandatory mTLS and disabled OAuth2 with an invalid client certificate in header
        Given The system is ready
        And I have a valid token with a client certificate "invalid-cert.txt"
        When I use the APK Conf file "artifacts/apk-confs/mtls/mtls_mandatory_oauth2_disabled.apk-conf"
        And the definition file "artifacts/definitions/employees_api.json"
        And make the API deployment request
        Then the response status code should be 200
        Then I set headers
            | X-WSO2-CLIENT-CERTIFICATE | ${clientCertificate} |
        And I send "GET" request to "https://default.gw.wso2.com:9095/mtls/3.14/employee/" with body ""
        And I eventually receive 401 response code, not accepting
            | 200 |
        When I undeploy the API whose ID is "mtls-mandatory-oauth2-disabled"
        Then the response status code should be 202

    Scenario: Test mandatory mTLS and disabled OAuth2 without a client certificate in header
        Given The system is ready
        And I have a valid subscription
        When I use the APK Conf file "artifacts/apk-confs/mtls/mtls_mandatory_oauth2_disabled.apk-conf"
        And the definition file "artifacts/definitions/employees_api.json"
        And make the API deployment request
        Then the response status code should be 200
        And I send "GET" request to "https://default.gw.wso2.com:9095/mtls/3.14/employee/" with body ""
        And I eventually receive 401 response code, not accepting
            | 200 |
        When I undeploy the API whose ID is "mtls-mandatory-oauth2-disabled"
        Then the response status code should be 202

    Scenario: Test optional mTLS and enabled OAuth2 with a valid client certificate in header
        Given The system is ready
        And I have a valid token with a client certificate "config-map-1.txt"
        When I use the APK Conf file "artifacts/apk-confs/mtls/mtls_optional_oauth2_enabled.apk-conf"
        And the definition file "artifacts/definitions/employees_api.json"
        And make the API deployment request
        Then the response status code should be 200
        Then I set headers
            | Authorization             | bearer ${accessToken} |
            | X-WSO2-CLIENT-CERTIFICATE | ${clientCertificate}  |
        And I send "GET" request to "https://default.gw.wso2.com:9095/mtls/3.14/employee/" with body ""
        And I eventually receive 200 response code, not accepting
            | 401 |
        When I undeploy the API whose ID is "mtls-optional-oauth2-enabled"
        Then the response status code should be 202

    Scenario: Test optional mTLS and enabled OAuth2 with an invalid client certificate in header
        Given The system is ready
        And I have a valid token with a client certificate "invalid-cert.txt"
        When I use the APK Conf file "artifacts/apk-confs/mtls/mtls_optional_oauth2_enabled.apk-conf"
        And the definition file "artifacts/definitions/employees_api.json"
        And make the API deployment request
        Then the response status code should be 200
        Then I set headers
            | Authorization             | bearer ${accessToken} |
            | X-WSO2-CLIENT-CERTIFICATE | ${clientCertificate}  |
        And I send "GET" request to "https://default.gw.wso2.com:9095/mtls/3.14/employee/" with body ""
        And I eventually receive 200 response code, not accepting
            | 401 |
        When I undeploy the API whose ID is "mtls-optional-oauth2-enabled"
        Then the response status code should be 202

    Scenario: Test optional mTLS and enabled OAuth2 without a client certificate in header
        Given The system is ready
        And I have a valid subscription
        When I use the APK Conf file "artifacts/apk-confs/mtls/mtls_optional_oauth2_enabled.apk-conf"
        And the definition file "artifacts/definitions/employees_api.json"
        And make the API deployment request
        Then the response status code should be 200
        Then I set headers
            | Authorization | bearer ${accessToken} |
        And I send "GET" request to "https://default.gw.wso2.com:9095/mtls/3.14/employee/" with body ""
        And I eventually receive 200 response code, not accepting
            | 401 |
        When I undeploy the API whose ID is "mtls-optional-oauth2-enabled"
        Then the response status code should be 202

    Scenario: Test optional mTLS and disabled OAuth2 with a valid client certificate in header
        Given The system is ready
        And I have a valid token with a client certificate "config-map-1.txt"
        When I use the APK Conf file "artifacts/apk-confs/mtls/mtls_optional_oauth2_disabled.apk-conf"
        And the definition file "artifacts/definitions/employees_api.json"
        And make the API deployment request
        Then the response status code should be 406

    Scenario: Test an API with mTLS disabled and OAuth2 disabled
        Given The system is ready
        And I have a valid subscription
        When I use the APK Conf file "artifacts/apk-confs/mtls/mtls_disabled_oauth2_disabled.apk-conf"
        And the definition file "artifacts/definitions/employees_api.json"
        And make the API deployment request
        Then the response status code should be 406

    Scenario: Test an API with mTLS disabled and OAuth2 enabled
        Given The system is ready
        And I have a valid subscription
        When I use the APK Conf file "artifacts/apk-confs/mtls/mtls_disabled_oauth2_enabled.apk-conf"
        And the definition file "artifacts/definitions/employees_api.json"
        And make the API deployment request
        Then the response status code should be 200
        Then I set headers
            | Authorization | bearer ${accessToken} |
        And I send "GET" request to "https://default.gw.wso2.com:9095/mtls/3.14/employee/" with body ""
        And I eventually receive 200 response code, not accepting
            | 401 |
        When I undeploy the API whose ID is "mtls-disabled-oauth2-enabled"
        Then the response status code should be 202

    Scenario: Test an API with mTLS enabled and one associated certificate with multiple certificates existing in system
        Given The system is ready
        And I have a valid token with a client certificate "config-map-1.txt"
        When I use the APK Conf file "artifacts/apk-confs/mtls/mtls_mandatory_oauth2_enabled.apk-conf"
        And the definition file "artifacts/definitions/employees_api.json"
        And make the API deployment request
        Then the response status code should be 200
        Then I set headers
            | Authorization             | bearer ${accessToken} |
            | X-WSO2-CLIENT-CERTIFICATE | ${clientCertificate}  |
        And I send "GET" request to "https://default.gw.wso2.com:9095/mtls/3.14/employee/" with body ""
        And I eventually receive 200 response code, not accepting
            | 401 |
        And I have a valid token with a client certificate "config-map-2.txt"
        Then I set headers
            | Authorization             | bearer ${accessToken} |
            | X-WSO2-CLIENT-CERTIFICATE | ${clientCertificate}  |
        And I send "GET" request to "https://default.gw.wso2.com:9095/mtls/3.14/employee/" with body ""
        And I eventually receive 401 response code, not accepting
            | 200 |
        And I have a valid token with a client certificate "config-map-3.txt"
        Then I set headers
            | Authorization             | bearer ${accessToken} |
            | X-WSO2-CLIENT-CERTIFICATE | ${clientCertificate}  |
        And I send "GET" request to "https://default.gw.wso2.com:9095/mtls/3.14/employee/" with body ""
        And I eventually receive 401 response code, not accepting
            | 200 |
        When I undeploy the API whose ID is "mtls-mandatory-oauth2-enabled"
        Then the response status code should be 202

    Scenario: Test an API with mTLS enabled and multiple certificates configured
        Given The system is ready
        And I have a valid token with a client certificate "config-map-1.txt"
        When I use the APK Conf file "artifacts/apk-confs/mtls/mtls_multiple_certs.apk-conf"
        And the definition file "artifacts/definitions/employees_api.json"
        And make the API deployment request
        Then the response status code should be 200
        Then I set headers
            | Authorization             | bearer ${accessToken} |
            | X-WSO2-CLIENT-CERTIFICATE | ${clientCertificate}  |
        And I send "GET" request to "https://default.gw.wso2.com:9095/mtls/3.14/employee/" with body ""
        And I eventually receive 200 response code, not accepting
            | 401 |
        And I have a valid token with a client certificate "config-map-2.txt"
        Then I set headers
            | Authorization             | bearer ${accessToken} |
            | X-WSO2-CLIENT-CERTIFICATE | ${clientCertificate}  |
        And I send "GET" request to "https://default.gw.wso2.com:9095/mtls/3.14/employee/" with body ""
        And I eventually receive 200 response code, not accepting
            | 401 |
        And I have a valid token with a client certificate "config-map-3.txt"
        Then I set headers
            | Authorization             | bearer ${accessToken} |
            | X-WSO2-CLIENT-CERTIFICATE | ${clientCertificate}  |
        And I send "GET" request to "https://default.gw.wso2.com:9095/mtls/3.14/employee/" with body ""
        And I eventually receive 401 response code, not accepting
            | 200 |
        When I undeploy the API whose ID is "mtls-multiple-certs"
        Then the response status code should be 202
