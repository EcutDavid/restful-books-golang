import React from 'react'
import request from 'superagent'

import { books } from '../config/apiUrl'
import 'styles/books.scss'

export default class Books extends React.Component {
  constructor() {
    super()
    this.state = { books: [] }
  }

  componentWillMount() {
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
          console.log(books);
          this.setState({ books })
        } catch (e) {
          return
        }
      });
  }

  render() {
    const { books } = this.state
    console.log(books);
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
                value={d.IsRent}
                className='input'
              />
            </div>
          ))
        }
      </div>
    )
  }
}
