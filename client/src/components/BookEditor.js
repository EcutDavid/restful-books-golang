import React from 'react'
import request from 'superagent'

import { books } from '../config/apiUrl'
import 'styles/bookEditor.scss'

export default class BookEditor extends React.Component {
  constructor() {
    super()
    this.state = { communicating: false }
  }

  addBook() {
    const { input: { value : bookName } } = this.refs
    const { bookSent } = this.props
    if (bookName === '') {
      alert('Please input name of the book ')
    }
    this.setState({ communicating: true })
    request
      .post(books)
      .type('form')
      .set('Accept', 'application/json')
      .send({ name: bookName })
      .end((err, res) => {
        if (err) {
          console.log(err);
          return
        }
        this.setState({ communicating: false })
        bookSent()
      })
  }

  render() {
    const { communicating } = this.state
    return (
      <div className='BookEditor'>
        <h3>Add new book</h3>
        <input
          type='text'
          placeholder='Book name'
          className='input'
          ref='input'
        />
        <button
          className='button'
          onClick={() => this.addBook()}
          disabled={communicating}
        >
          send
        </button>
      </div>
    )
  }
}
