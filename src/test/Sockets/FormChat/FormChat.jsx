import { Button, Col, Divider, Form, Input, Row, Space } from 'antd'
import React from 'react'
import './form.css'

const { TextArea } = Input;

export default function FormChat() {
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
