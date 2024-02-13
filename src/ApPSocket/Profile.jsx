import React, { useEffect, useState } from 'react'
import UserContext from './UserContext';
import { Avatar, Button, Input, message } from 'antd';
import axios from 'axios';
import { SlackOutlined } from '@ant-design/icons';

export default function Profile({ backToHome }) {

  const [information, setInformation] = useState();
  const [informationFindByUserName, setInformationFindByUserName] = useState('');

  const userName = UserContext();



  const handleLogout = () => {
    localStorage.removeItem('user_name');
    backToHome();
  };

  useEffect(() => {

    const getProfileByUserName = async () => {
      try {
        const response = await axios.get(`http://localhost:8080/manager/find?user_name=${userName}`);
        if (response.data.result.code === 0) {
          setInformation(response.data.user);
          console.log(response.data.user);
        } else {
          message.error("lỗi server 1");
        }
      } catch (error) {
        message.error("lỗi server");
      }
    }
    if (userName) {
      getProfileByUserName();
    }
  }, [userName]);


  const handleFindUserByUserName = async () => {
  
    if (informationFindByUserName.trim() === '') {
      message.warning('Vui lòng nhập thông tin');
      return;
    }

    try {
      const response = await axios.get(`http://localhost:8080/manager/find?user_name=${informationFindByUserName}`);
      if (response.data.result.code === 0) {
        setInformationFindByUserName(response.data.user);
      } else if (response.data.result.code === 7) {
        setInformationFindByUserName();
        message.warning('Không tìn thấy user');
      } else {
        message.error('Lỗi server');
      }
    } catch (error) {
      message.error('Lỗi server');
    }
  };

  const handleChange = (e) => {
    const value = e.target.value;
   
    setInformationFindByUserName(value);

  };

  return (
    <div>
      <div>

        <ul style={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between' }}>

          <ul style={{ display: 'flex', listStyle: 'none' }}>

            <li style={{ width: '60px' }}>
              <SlackOutlined />
            </li>

            <li style={{ width: '120px' }}>
              Chào mừng bạn
            </li>

          </ul>

          <ul style={{ display: 'flex', alignItems: 'center', marginRight: '100px', listStyle: 'none' }}>

            <li style={{ width: '90px', fontSize: '20px' }}>

              {information && information.user_name}
            </li>

            <li style={{ width: '60px' }}>
              {information && <Avatar src={information.avatar} />}
            </li>

            <li style={{ width: '60px' }}>
              <Button onClick={handleLogout}>Logout</Button>
            </li>

          </ul>

        </ul>

      </div>

      <div>
        <ul style={{ display: 'flex', listStyle: 'none' }}>
          <li>
            <Input
              required
              onChange={handleChange} />
          </li>
          <li>
            <Button onClick={handleFindUserByUserName}>Tìm kiếm</Button>
          </li>
        </ul>

        {informationFindByUserName && (

          <ul style={{ listStyle: 'none' }}>
            <li>{informationFindByUserName.user_name}</li>
            <li>{informationFindByUserName.email}</li>
            <li>{informationFindByUserName.address}</li>
            <li>{informationFindByUserName.age}</li>
            <li>
              <Avatar src={informationFindByUserName.avatar}/>
            </li>
          </ul>
        )}
      </div>

    </div>
  )
}
