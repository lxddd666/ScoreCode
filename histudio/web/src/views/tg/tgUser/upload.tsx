import  { useState } from 'react';
import Button from '@mui/material/Button';

export default function FileUpload() {
  const [selectedFile, setSelectedFile] = useState(null);

  const handleFileChange = (event:any) => {
    const file = event.target.files[0];
    // if(file.size / 1024 / 1024 > 2){
    //     alert('文件大小不能大于 2MB');
    //     return
    // }
    if (file && file.name.endsWith('.xlsx')) {
      setSelectedFile(file);
    } else {
      alert('Please select an .xlsx file.');
      setSelectedFile(null);
    }
  };

  const handleUpload = () => {
    if (!selectedFile) {
      alert('Please select a file to upload.');
      return;
    }

    const formData = new FormData();
    formData.append('file', selectedFile);

    // 在这里你可以发送一个请求到后端来上传文件
    // 例如，使用 fetch 或 axios：
    // fetch('YOUR_BACKEND_ENDPOINT', {
    //   method: 'POST',
    //   body: formData,
    // }).then(response => {
    //   // 处理响应
    // });

    // alert('File uploaded successfully.'); // 模拟上传成功
  };

  return (
    <div>
      <input
        accept=".xlsx"
        style={{ display: 'none' }}
        id="raised-button-file"
        type="file"
        onChange={handleFileChange}
      />
      <label htmlFor="raised-button-file">
        <Button variant="contained" component="span">
          Select .xlsx File
        </Button>
      </label>
      <Button variant="contained" color="primary" onClick={handleUpload} disabled={!selectedFile}>
        Upload
      </Button>
    </div>
  );
}
