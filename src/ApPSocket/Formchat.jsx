import React from 'react'
import './index.css'
import { Button, Input, Space } from 'antd'
const { TextArea } = Input;

export default function Formchat() {

    

    return (
        <div>
          <div className='form-container-message'>
            <TextArea
              style={{ backgroundColor: 'white' }}
              disabled
              placeholder="Autosize height with minimum and maximum number of lines"
              autoSize={{
                minRows: 30,
                maxRows: 6,
              }}
            />
    
            <div style={{ display: 'flex' }}>
              <Input style={{ borderRadius: '0' }} />
              <Button style={{ borderRadius: '0' }}>Send</Button>
            </div>
    
          </div>
        </div>
      )
}

