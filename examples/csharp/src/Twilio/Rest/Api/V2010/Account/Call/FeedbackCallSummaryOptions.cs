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

using Twilio.Types;


namespace Twilio.Rest.Api.V2010.Account.Call
{

    public class CreateFeedbackCallSummaryOptions : IOptions<FeedbackCallSummaryResource>
    {
        public DateTime? EndDate { get; set; }
        public DateTime? StartDate { get; set; }
        public string PathAccountSid { get; set; }

        public CreateFeedbackCallSummaryOptions(DateTime? endDate, DateTime? startDate)
        {
            EndDate = endDate;
            StartDate = startDate;
        }

        
        public  List<KeyValuePair<string, string>> GetParams()
        {
            var p = new List<KeyValuePair<string, string>>();

            if (EndDate != null)
            {
                p.Add(new KeyValuePair<string, string>("EndDate", EndDate.Value.ToString("yyyy-MM-dd")));
            }
            if (StartDate != null)
            {
                p.Add(new KeyValuePair<string, string>("StartDate", StartDate.Value.ToString("yyyy-MM-dd")));
            }
            return p;
        }
        

    }
}

