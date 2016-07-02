import React from 'react'
import request from 'superagent'

import BookEditor from './BookEditor'
import { books } from '../config/apiUrl'
import 'styles/books.scss'

export default class Books extends React.Component {
  constructor() {
    super()
    this.state = { books: [] }
  }

  componentWillMount() {
    this.updateBooks()
  }

  updateBooks() {
    request
      .get(books)
      .set('Accept', 'application/json')
      .send()
      .end((err, res) => {
        if (err) {
          console.log(err);
          return
        }
        try {
          const books = JSON.parse(res.text)
          this.setState({ books })
        } catch (e) {
          return
        }
      })
  }

  rentBook(id, isRent) {
    request
      .put(`${books}/${id}`)
      .set('Accept', 'application/json')
      .send({ isRent: isRent ? 'false' : 'true' })
      .type('form')
      .end((err, res) => {
        if (err) {
          console.log(err);
          return
        }
        this.updateBooks()
      })
  }

  render() {
    const { books } = this.state
    return (
      <div className='Books'>
        <div className='book-area header'>
          <p className='name'>
            Book Name
          </p>
          <p className='input'>
            Is Rent
          </p>
        </div>
        {
          books.map((d, key) => (
            <div className='book-area' key={key}>
              <p className='name'>
                { d.Name }
              </p>
              <input
                type='checkbox'
                defaultChecked={d.IsRent}
                className='input'
                onClick={() => this.rentBook(d.Id, d.IsRent)}
              />
            </div>
          ))
        }
        <BookEditor bookSent={() => this.updateBooks()} />
      </div>
    )
  }
}
