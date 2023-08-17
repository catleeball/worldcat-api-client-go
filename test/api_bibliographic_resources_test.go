/*
WorldCat Search API v. 2

Testing BibliographicResourcesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
    "context"
    openapiclient "github.com/catleeball/worldcat-api-client-go"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
)

func Test_openapi_BibliographicResourcesApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test BibliographicResourcesApiService RetrieveBib", func(t *testing.T) {

        t.Skip("skip test") // remove to run test

        var oclcNumber int64

        resp, httpRes, err := apiClient.BibliographicResourcesApi.RetrieveBib(context.Background(), oclcNumber).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BibliographicResourcesApiService RetrieveBriefBib", func(t *testing.T) {

        t.Skip("skip test") // remove to run test

        var oclcNumber int64

        resp, httpRes, err := apiClient.BibliographicResourcesApi.RetrieveBriefBib(context.Background(), oclcNumber).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BibliographicResourcesApiService RetrieveOtherEditions", func(t *testing.T) {

        t.Skip("skip test") // remove to run test

        var oclcNumber int64

        resp, httpRes, err := apiClient.BibliographicResourcesApi.RetrieveOtherEditions(context.Background(), oclcNumber).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BibliographicResourcesApiService SearchBibs", func(t *testing.T) {

        t.Skip("skip test") // remove to run test

        resp, httpRes, err := apiClient.BibliographicResourcesApi.SearchBibs(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BibliographicResourcesApiService SearchBriefBibs", func(t *testing.T) {

        t.Skip("skip test") // remove to run test

        resp, httpRes, err := apiClient.BibliographicResourcesApi.SearchBriefBibs(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
