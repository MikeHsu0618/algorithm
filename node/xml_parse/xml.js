fs = require('fs');
var parser = require('xml2json');

fs.readFile( './example.xml', function(err, data) {
    var json = JSON.parse(parser.toJson(data));
    let competitor = []
    let inningnr = []
    json['market_descriptions']['market'].forEach(item => {
        if (item.name.includes("{$competitor2}")) {
            competitor.push(item)
            return
        }
        if (item.name.includes("inningnr")) {
            inningnr.push(item)
        }
    })

    console.log(JSON.stringify(competitor, null, 4))
    console.log("length: ", competitor.length)

    // console.log(JSON.stringify(inningnr, null, 4))
});