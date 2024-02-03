import { SlackOutlined } from '@ant-design/icons';
import { Button, Form, Input, Modal, Upload, message } from 'antd';
import React, { useState } from 'react'
import './index.css'
import axios from 'axios';
export default function FormSignUp() {


  const layout = {
    labelCol: {
      span: 8,
    },
    wrapperCol: {
      span: 16,
    },
    wrapperCol: {
      offset: 0,
    },
  };

  const [modal1Open, setModal1Open] = useState(false);


  const errorMessage = () => {
    message.error('Lỗi hệ thống vui lòng thử lại');
  };

  const sucessMessage = () => {
    message.success("Tạo tài khoản thành công");
  }
  const warningMesage = () => {
    message.warning("Tài khoản  tồn tại");
  }

  const [imageFile, setImageFile] = useState(null);

  const handleFormSubmit = async (values) => {
    try {
      //const [] = Form.useForm();
      const formData = new FormData();
      
      // Upload file
      formData.append('file', imageFile);

      // Set form data
      formData.append('user_name', values.user_name);
      formData.append('age', values.age);
      formData.append('address', values.address);
      formData.append('email', values.email);

      // Gửi dữ liệu lên máy chủ sử dụng Axios
      const response = await axios.post('http://localhost:8080/manager/add', formData);

      // Xử lý kết quả từ máy chủ
      if (response.data.result.code === 5) {
        warningMesage();
      } else if (response.data.result.code === 1) {
        errorMessage();
      } else {
        sucessMessage();
      }

    } catch (error) {
      console.error(error);
      alert('Có lỗi xảy ra. Vui lòng thử lại.');
      return;
    }
  };


  return (
    <div>
      <Button type="primary" onClick={() => setModal1Open(true)}>
        Đăng ký tài khoản
      </Button>
      <Modal

        footer
        style={{
          top: 20,
        }}
        open={modal1Open}
        onOk={() => setModal1Open(false)}
        onCancel={() => setModal1Open(false)}
      >
        <Form
          {...layout}
          className='sign-up-form'
          onFinish={handleFormSubmit}
        >
          <Form.Item>
            <div style={{ fontSize: '30px', color: 'gray' }}>
              Sign up
              <SlackOutlined />
            </div>
          </Form.Item>
          <div className='login-form-home-element'>

            <Form.Item
              name='user_name'
              label='Tên tài khoản'
            >
              <Input className='form-login-input' />
            </Form.Item>

            <Form.Item
              name='age'
              label='Tuổi'
            >
              <Input className='form-login-input' />
            </Form.Item>

            <Form.Item
              name='address'
              label='Địa chỉ'
            >
              <Input className='form-login-input' />
            </Form.Item>

            <Form.Item
              name='email'
              label='Email'
            >
              <Input className='form-login-input' />
            </Form.Item>

            <Form.Item
              label='Avatar'
              name='file'
              valuePropName="fileList"
              getValueFromEvent={(e) => {
                if (Array.isArray(e)) {
                  return e;
                }
                return e && e.fileList;
              }}
            >
              <Upload
                maxCount={1}
                type=''
                listType='picture-card'
                openFileDialogOnClick
                accept="image/jpeg,image/png"
                beforeUpload={(file) => {
                  setImageFile(file);
                  return false;
                }}
                onRemove={() => {
                  setImageFile(null);
                }}
              >
                +Upload

              </Upload>
            </Form.Item>

            <Form.Item
            >
              <Button htmlType='submit' className='button-login-app-chat'>
                Sign up</Button>
            </Form.Item>

          </div>
        </Form>
      </Modal>

    </div>
  )
}
