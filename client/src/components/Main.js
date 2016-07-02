import 'normalize.css/normalize.css'
import 'styles/app.scss'
import React from 'react'

import Header from './Header'
import Books from './Books'

class AppComponent extends React.Component {
  render() {
    return (
      <div className="Main">
        <Header />
        <Books />
      </div>
    )
  }
}

AppComponent.defaultProps = {
}

export default AppComponent
