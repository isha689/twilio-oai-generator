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


namespace Twilio.Rest.Api.V2010.Credential
{

    public class DeleteAwsOptions : IOptions<AwsResource>
    {
        
        public string PathSid { get; set; }


        /// <summary>
        /// Construct a new DeleteCredentialAwsOptions
        /// </summary>
        /// <param name="pathSid">  </param>

        public DeleteAwsOptions(string pathSid)
        {
            PathSid = pathSid;
        }
    }



    public class FetchAwsOptions : IOptions<AwsResource>
    {
    
        public string PathSid { get; set; }


        /// <summary>
        /// Construct a new FetchCredentialAwsOptions
        /// </summary>
        /// <param name="pathSid">  </param>

        public FetchAwsOptions(string pathSid)
        {
            PathSid = pathSid;
        }
    }



    public class ReadAwsOptions : ReadOptions<AwsResource>
    {
    
        public int? PageSize { get; set; }


    }


    public class UpdateAwsOptions : IOptions<AwsResource>
    {
    
        public string PathSid { get; set; }
        public string TestString { get; set; }
        public bool? TestBoolean { get; set; }


        /// <summary>
        /// Construct a new UpdateCredentialAwsOptions
        /// </summary>
        /// <param name="pathSid">  </param>

        public UpdateAwsOptions(string pathSid)
        {
            PathSid = pathSid;
        }
    }


}
