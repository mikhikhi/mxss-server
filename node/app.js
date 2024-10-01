const sanitizeHtml = require('sanitize-html');
const express = require('express');
const createDOMPurify = require('dompurify');
const { JSDOM } = require('jsdom');

const window = new JSDOM('').window;
const DOMPurify = createDOMPurify(window);

app = express()

app.use(express.urlencoded({ extended: false }));


const PORT = process.env.PORT || 3000;

app.post('/', (req, res) => {
  res.status(200).send(sanitizeHtml(req.body.js));
});


app.post('/dompurify', (req, res) => {
  res.status(200).send(DOMPurify.sanitize(req.body.js));
});

app.listen(PORT, () => {
  console.log(`Server is listening on port ${PORT}`);
});