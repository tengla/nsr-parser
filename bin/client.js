const https = require('http');

const prettyPrintPlace = place => {
  console.log(
    place.id, place.name.text,
    '(' + place.TransportMode + ')',
    place.keylist.keyvalue
      .filter(kv => kv.Key.match(/^(jbv|uic|grails)(Code|Id)$/))
      .reduce((o, c) => {
        return {
          [c.Key]: c.Value,
          ...o
        }
      }, place.centroid.location)
  );
};

const httpsAsync = (query) => {
  return new Promise((resolve, reject) => {
    const req = https.request({
      // hostname: "nsr-o5iqmbtjaa-ew.a.run.app",
      hostname: "127.0.0.1",
      port: "8080",
      method: "POST",
      headers: {
        'Content-Type': 'application/json',
        'Content-Length': query.length
      }
    }, res => {
      if (res.statusCode > 299) {
        reject({
          statusCode: res.statusCode,
          statusMessage: res.statusMessage
        })
      }
      const chunks = [];
      res.on('data', chunk => chunks.push(chunk))
      res.on('end', () => {
        const docs = Buffer.concat(chunks).toString().split("\n").map(doc => {
          if (doc.length > 0) {
            return JSON.parse(doc);
          }
        });
        resolve(docs);
      })
    })
    req.on('error', err => {
      // never mind conn reset
      //if (err.code != 'ECONNRESET') {
        console.error(err);
      //}
    });
    req.write(query);
    req.end()
  });
}
const data = Buffer.from(JSON.stringify({
  "latitude": 59.90724768268159,
  "longitude": 10.753830146286921,
  "distance": 15,
  "name": "oslo",
  //"transportmode":"rail",
  //"key": "jbvCode",
  //"value": "LOD"
}), 'utf-8');

httpsAsync(data)
  .then(places => {
    places
      .filter(p => {
        return p && Object.prototype.hasOwnProperty.call(p, "id")
      })
      .forEach(prettyPrintPlace)
  })
  .catch(console.error);
