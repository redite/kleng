import React, { Component } from 'react'
import Header from './Header'
import Sidebar from './Sidebar'
import Switcher from './Switcher'
import { Row } from '../theme/globalStyle'

class App extends Component {
  state = {
    currentPage: 'home',
    currentLang: '',
    currentLabel: '',
  }

  onLanguageChange = lang => e => {
    e.preventDefault()
    this.setState({
      currentLang: lang,
    })
  }

  onLabelChange = label => e => {
    e.preventDefault()
    this.setState({
      currentLabel: label,
    })
  }

  render() {
    return (
      <div>
        <Header />
        <div className="container-fluid">
          <Row>
            <Sidebar
              onLanguageChange={this.onLanguageChange}
              onLabelChange={this.onLabelChange}
            />
            <Switcher />
          </Row>
        </div>
      </div>
    )
  }
}

export default App
