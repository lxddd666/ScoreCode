// useConfirm.js
import ReactDOM from 'react-dom';

import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import Button from '@mui/material/Button';

const ConfirmDialog = ({ open, onClose, onConfirm, title, content }: any) => {
  return (
    <Dialog open={open} onClose={onClose}>
      <DialogTitle>{title}</DialogTitle>
      <DialogContent>
        <DialogContentText>{content}</DialogContentText>
      </DialogContent>
      <DialogActions>
        <Button onClick={() => onClose(false)} color="primary">
          取消
        </Button>
        <Button onClick={() => onConfirm(true)} color="primary" autoFocus>
          确定
        </Button>
      </DialogActions>
    </Dialog>
  );
};

const useConfirm = () => {
  const confirm = (title: any, content: any) => {
    return new Promise((resolve, reject) => {
      const handleClose = (result: any) => {
        cleanup();
        resolve(result);
      };

      const handleConfirm = (result: any) => {
        cleanup();
        resolve(result);
      };

      const cleanup = () => {
        const div: any = document.getElementById('confirm-dialog-container');
        ReactDOM.unmountComponentAtNode(div);
        div.parentNode.removeChild(div);
      };

      const div = document.createElement('div');
      div.id = 'confirm-dialog-container';
      document.body.appendChild(div);

      ReactDOM.render(
        <ConfirmDialog
          open={true}
          title={title}
          content={content}
          onClose={handleClose}
          onConfirm={handleConfirm}
        />,
        div
      );
    });
  };

  return confirm;
};


export default useConfirm;
