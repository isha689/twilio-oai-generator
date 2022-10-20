/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Versionless
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using Twilio.Base;
using Twilio.Clients;
using Twilio.Converters;
using Twilio.Exceptions;
using Twilio.Http;



namespace Twilio.Rest.Versionless.Understand
{
    public class AssistantResource : Resource
    {
    

        
        private static Request BuildCreateRequest(CreateAssistantOptions options, ITwilioRestClient client)
        {
            
            string path = "/understand/Assistants";



            return new Request(
                HttpMethod.Post,
                Rest.Domain.Versionless,
                path,
                postParams: options.GetParams(),
                headerParams: null
            );
        }

        /// <summary> create </summary>
        /// <param name="options"> Create Assistant parameters </param>
        /// <param name="client"> Client to make requests to Twilio </param>
        /// <returns> A single instance of Assistant </returns>
        public static AssistantResource Create(CreateAssistantOptions options, ITwilioRestClient client = null)
        {
            client = client ?? TwilioClient.GetRestClient();
            var response = client.Request(BuildCreateRequest(options, client));
            return FromJson(response.Content);
        }

        #if !NET35
        /// <summary> create </summary>
        /// <param name="options"> Create Assistant parameters </param>
        /// <param name="client"> Client to make requests to Twilio </param>
        /// <returns> Task that resolves to A single instance of Assistant </returns>
        public static async System.Threading.Tasks.Task<AssistantResource> CreateAsync(CreateAssistantOptions options,
        ITwilioRestClient client = null)
        {
            client = client ?? TwilioClient.GetRestClient();
            var response = await client.RequestAsync(BuildCreateRequest(options, client));
            return FromJson(response.Content);
        }
        #endif

        /// <summary> create </summary>
        /// <param name="friendlyName">  </param>
        /// <param name="client"> Client to make requests to Twilio </param>
        /// <returns> A single instance of Assistant </returns>
        public static AssistantResource Create(
                                          string friendlyName = null,
                                          ITwilioRestClient client = null)
        {
            var options = new CreateAssistantOptions(){  FriendlyName = friendlyName };
            return Create(options, client);
        }

        #if !NET35
        /// <summary> create </summary>
        /// <param name="friendlyName">  </param>
        /// <param name="client"> Client to make requests to Twilio </param>
        /// <returns> Task that resolves to A single instance of Assistant </returns>
        public static async System.Threading.Tasks.Task<AssistantResource> CreateAsync(
                                                                                  string friendlyName = null,
                                                                                  ITwilioRestClient client = null)
        {
        var options = new CreateAssistantOptions(){  FriendlyName = friendlyName };
            return await CreateAsync(options, client);
        }
        #endif




    
        /// <summary>
        /// Converts a JSON string into a AssistantResource object
        /// </summary>
        /// <param name="json"> Raw JSON string </param>
        /// <returns> AssistantResource object represented by the provided JSON </returns>
        public static AssistantResource FromJson(string json)
        {
            try
            {
                return JsonConvert.DeserializeObject<AssistantResource>(json);
            }
            catch (JsonException e)
            {
                throw new ApiException(e.Message, e);
            }
        }

    
        ///<summary> The unique SID that identifies this Account. </summary> 
        [JsonProperty("account_sid")]
        public string AccountSid { get; private set; }

        ///<summary> A human readable description for this Fleet. </summary> 
        [JsonProperty("friendly_name")]
        public string FriendlyName { get; private set; }

        ///<summary> A string that uniquely identifies this Fleet. </summary> 
        [JsonProperty("sid")]
        public string Sid { get; private set; }



        private AssistantResource() {

        }
    }
}

