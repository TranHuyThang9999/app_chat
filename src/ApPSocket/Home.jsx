import React from 'react'
import './index.css'
import { MessageOutlined } from '@ant-design/icons';
import FormLogin from './FormSignIn';
import FormSignUp from './FormSignUp';

export default function Home({onLoginSuccess}) {

  return (
    <div className='container-home'>
      <ul className='nav-home'>
        <li className='content-start-home'>
          <div style={{ marginRight: '9px' }}>
            Chào mừng bạn
          </div>
          <div className='message-home'>
            <MessageOutlined />
          </div>
        </li>
        <li>
          <FormSignUp />
        </li>
      </ul>
      <div >
        <FormLogin onLoginSuccess={onLoginSuccess} />
      </div>
    </div>
  )
}

