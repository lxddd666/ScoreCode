import { useState } from 'react';
import Button from '@mui/material/Button';
import axios from 'axios'

const FileUpload = (props:any)=> {
  const {selectedFileChange} = props
  const [selectedFile, setSelectedFile] = useState(null);

  const handleFileChange = (event: any) => {
    const file = event.target.files[0];
    // if(file.size / 1024 / 1024 > 2){
    //     alert('文件大小不能大于 2MB');
    //     return
    // }

    if (file && file.name.endsWith('.zip')) {
      selectedFileChange(file)
      setSelectedFile(file);
    } else {
      alert('只能上传zip格式的文件，请重新上传');
      setSelectedFile(null);
    }
    setSelectedFile(file);
  };

  const handleUpload = (e: any) => {
    console.log('1', e)
    if (!selectedFile) {
      alert('只能上传zip格式的文件，请重新上传');
      return;
    }

    const formData = new FormData();
    formData.append('file', selectedFile);

    // 在这里你可以发送一个请求到后端来上传文件
    // 例如，使用 fetch 或 axios：
    axios('http://10.8.12.88:8001/tg/tgUser/importSession', {
      method: 'POST',
      transformRequest: [function (data, headers:any) {
        // 去除post请求默认的Content-Type
        // console.log(data, headers);
        // delete headers.post['Content-Type']
        return data
      }],
      data: formData,
    }).then(res => {
      // 处理响应
      console.log('res上传成功', res);

    }).catch(err => {
      console.log('res上传失败', err);
    })

    // alert('File uploaded successfully.'); // 模拟上传成功
  };

  return (
    <div>
      <input
        style={{ display: 'none' }}
        accept=".zip,"
        id="raised-button-file"
        type="file"
        onChange={handleFileChange}
      />
      <label htmlFor="raised-button-file">
        <Button variant="contained" component="span">
          选择文件
        </Button>
      </label>
      <Button variant="contained" color="primary" onClick={handleUpload} disabled={!selectedFile}>
        {/* Upload */}
        {selectedFile ? '已上传':'请上传文件'}
      </Button>
    </div>
  );
}

export default FileUpload