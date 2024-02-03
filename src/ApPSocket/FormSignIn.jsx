import { DingdingOutlined } from '@ant-design/icons';
import { Button, Form, Input, message } from 'antd';
import axios from 'axios';
import React, { useContext } from 'react';

export default function FormLogin({onLoginSuccess}) {


  const tailLayout = {
    wrapperCol: {
      offset: 0,
    },
  };

  const errorMessage = () => {
    message.error('Lỗi hệ thống vui lòng thử lại');
  };

  const successMessage = () => {
    message.success('Đăng nhập thành công');
  };


  const handleLogin = async (values) => {
    try {


      const formData = new FormData();
      formData.append('user_name', values.user_name);
      formData.append('email', values.email);

      const response = await axios.post('http://localhost:8080/manager/login', formData);

      if (response.data.result.code === 0) {

          localStorage.setItem('user_name',values.user_name);
        onLoginSuccess();
        successMessage();
      } else {
        message.warning('Thông tin tài khoản hoặc mật khẩu không chính xác. Vui lòng thử lại.');
      }

    } catch (error) {
      errorMessage();
      console.error(error);
    }
  };

  return (
    <div>
      <Form
        {...tailLayout}
        className="login-form"
        onFinish={handleLogin}
      >
        <Form.Item>
          <div style={{ fontSize: '30px', color: 'gray' }}>
            Sign In
            <DingdingOutlined />
          </div>
        </Form.Item>
        <div className="login-form-home-element">
          <Form.Item
            rules={[
              {
                required: true,
                message: 'Vui lòng nhập tên tài khoản',
                whitespace: true,
              },
            ]}
            name="user_name"
          >
            <Input
              name="user_name"
              className="form-login-input"
              placeholder="Nhập tên tài khoản"
            />
          </Form.Item>
          <Form.Item
            rules={[
              {
                required: true,
                message: 'Vui lòng nhập địa chỉ email',
                whitespace: true,
              },
            ]}
            name="email"
            className="form-login-input"
          >
            <Input
              name="email"
              className="form-login-input"
              placeholder="Nhập địa chỉ email"
            />
          </Form.Item>
          <Form.Item>
            <Button  htmlType="submit" className="button-login-app-chat">
              Login
            </Button>
          </Form.Item>
        </div>
      </Form>
    </div>
  );
}