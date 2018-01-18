pragma solidity 0.4;


contract DarkID {

    struct IdStruct {
        string pubK;
        string date;
        string hashed;
        string unblindedSig;
        string serverVerifier;
        string signerID;
    }


    IdStruct public ID;


    function DarkID(string _pubK, string _hashed, string _unblindedSig, string _serverVerifier, string _signerID) public {
        ID = IdStruct(_pubK, "this will be the date", _hashed, _unblindedSig, _serverVerifier, _signerID);
    }


    function getDarkID() public constant returns(string, string, string, string, string, string) {
        return (ID.pubK, ID.date, ID.hashed, ID.unblindedSig, ID.serverVerifier, ID.signerID);
    }
}
