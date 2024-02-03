import React from 'react'
import UserContext from './UserContext';
import { Button } from 'antd';

export default function Profile() {

  const userName = UserContext();
  if (userName == null) {

  }

  return (
    <div>
      <div>
        <div>
     
        </div>
        <h2>userName :  {userName}</h2>

      </div>
    </div>
  )
}
