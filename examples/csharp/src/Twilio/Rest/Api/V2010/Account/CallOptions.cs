/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Accounts
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


using System;
using System.Collections.Generic;
using Twilio.Base;
using Twilio.Converters;
using System.Linq;



namespace Twilio.Rest.Api.V2010.Account
{

    public class CreateCallOptions : IOptions<CallResource>
    {
        // How to decide which has getter and which has setter ?
        
        public string RequiredStringProperty { get; }
        public string PathAccountSid { get; set; }
        public List<string> TestArrayOfStrings { get; set; }
        public List<Uri> TestArrayOfUri { get; set; }

        public CreateCallOptions(string requiredStringProperty)
        {
            RequiredStringProperty = requiredStringProperty;
            TestArrayOfStrings = new List<string>();
            TestArrayOfUri = new List<Uri>();
        }

        
        public  List<KeyValuePair<string, string>> GetParams()
        {
            var p = new List<KeyValuePair<string, string>>();

            if (RequiredStringProperty != null)
            {
                p.Add(new KeyValuePair<string, string>("RequiredStringProperty", RequiredStringProperty));
            }
            if (TestArrayOfStrings != null)
            {
                p.AddRange(TestArrayOfStrings.Select(TestArrayOfStrings => new KeyValuePair<string, string>("TestArrayOfStrings", TestArrayOfStrings)));
            }
            if (TestArrayOfUri != null)
            {
                p.AddRange(TestArrayOfUri.Select(TestArrayOfUri => new KeyValuePair<string, string>("TestArrayOfUri", Serializers.Url(TestArrayOfUri))));
            }
            return p;
        }
        

    }
    public class DeleteCallOptions : IOptions<CallResource>
    {
        
        public int? PathTestInteger { get; }
        public string PathAccountSid { get; set; }


        public DeleteCallOptions(int? pathTestInteger)
        {
            PathTestInteger = pathTestInteger;
        }

        
        public  List<KeyValuePair<string, string>> GetParams()
        {
            var p = new List<KeyValuePair<string, string>>();

            return p;
        }
        

    }


    public class FetchCallOptions : IOptions<CallResource>
    {
    
        public int? PathTestInteger { get; }
        public string PathAccountSid { get; set; }


        public FetchCallOptions(int? pathTestInteger)
        {
            PathTestInteger = pathTestInteger;
        }

        
        public  List<KeyValuePair<string, string>> GetParams()
        {
            var p = new List<KeyValuePair<string, string>>();

            return p;
        }
        

    }


}
