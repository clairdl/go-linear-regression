const csv = require('csv-parse');
const createCsvWriter = require('csv-writer').createArrayCsvWriter;

const fs = require('fs');
const outPath = '../iris_setosa.csv';

let res = [];

fs.createReadStream('../iris_headers.csv')
  .pipe(csv())
  .on('data', (row) => {
    if (row.pop() === 'Iris-setosa') {
      res.push(row);
      console.log(row);
    }
  })
  .on('error', (err) => console.log('Error: ', err))
  .on('end', () => {
    csvWriter.writeRecords(res)
      .then(() => {
        console.log('Successfully written to:', outPath);
      });
});

const csvWriter = createCsvWriter({
  header: ['Sepal length', 'Sepal width', 'Petal length', 'Petal width'],
  path: outPath,
});

