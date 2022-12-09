v1.0.8
Modified updateMatruingPrice function as follows:
function updateMaturingPrice( uint256 _predictionPrice, bool _samePredictionPrizeToken, uint256 _prizePrice, address _oracleContract, uint256 _oracleDecimals ) external onlyPoolCreator {}

Important notes: 

If _samePredictionPrizeToken is true, the following 3 params are ignored, and based on the oracleContract address (zero or non-zero), the price will be set by chainlink or by the  _predictionPrice.

If _samePredictionPrizeToken is false, either a valid _oracleContract address must be specified and _prizePrice and _oracleDecimals are ignored and should be left zero, or if an oracle is not available for prize, _prizePrice and _oracleDecimals must be set correctly.

Notice:

Upon pool creation, _wrappedTokenContract address must be the address of the prize token and _wrappedTokenSymbol is the symbol of the prediction token