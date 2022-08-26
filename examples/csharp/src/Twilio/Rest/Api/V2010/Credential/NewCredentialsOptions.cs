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

    public class CreateNewCredentialsOptions : IOptions<NewCredentialsResource>
    {
        // How to decide which has getter and which has setter ?
        
        public string TestString { get; set; }
        public bool? TestBoolean { get; set; }
        public int? TestInteger { get; set; }
        public decimal? TestNumber { get; set; }
        public float? TestNumberFloat { get; set; }
        public double? TestNumberDouble { get; set; }
        public decimal? TestNumberInt32 { get; set; }
        public long? TestNumberInt64 { get; set; }
        public Object TestObject { get; set; }
        public DateTime? TestDateTime { get; set; }
        public DateTime? TestDate { get; set; }
        public TestResource.StatusEnum TestEnum { get; set; }
        public List<Object> TestObjectArray { get; set; }
        public object TestAnyType { get; set; }
        ///<summary>
        /// A comma-separated list of the permissions you will request from the users of this ConnectApp.  Can include: &#x60;get-all&#x60; and &#x60;post-all&#x60;.
        ///</summary>
        public List<string> Permissions { get; set; }

        /// <summary>
        /// Construct a new CreateCredentialAwsOptions
        /// </summary>
        /// <param name="testString">  </param>

        public CreateNewCredentialsOptions(string testString)
        {
            TestString = testString;
            TestObjectArray = new List<Object>();
            Permissions = new List<string>();
        }
    }
}

