const express = require('express');
const app = express();
const port = process.env.PORT || 3000;

app.get('/', (req, res) => {
  res.json({
    message: 'مرحباً بك في تطبيق غيمة التجريبي!',
    timestamp: new Date().toISOString()
  });
});

app.listen(port, () => {
  console.log(`التطبيق يعمل على المنفذ ${port}`);
});
