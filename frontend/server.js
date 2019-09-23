let express = require('express')
let path = require('path')

const port = 3000

let app = express()

app.get('/', function (req, res) {
    res.sendFile('index.html', { root: path.join(__dirname, '../') })
})

app.use('/static', express.static('static'))

app.listen(port);

console.log(`server is listening on ${port}`)