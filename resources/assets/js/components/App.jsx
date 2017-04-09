import React, { Component } from 'react';
import Nav from './Nav';

class App extends Component {
  render() {
    return (
      <div>
        <Nav />
        <div className="tm-sidebar-left">
          <h3>Repo</h3>
          <ul className="uk-nav uk-nav-default tm-nav">
            <li className="uk-nav-header">PHP</li>
            <li><a>sonnet</a></li>
            <li><a>react-native</a></li>
            <li><a>alm</a></li>
            <li><a>beacon-ml</a></li>
          </ul>
          <ul className="uk-nav uk-nav-default tm-nav uk-margin-top">
            <li className="uk-nav-header">Go</li>
            <li><a>sonnet</a></li>
            <li><a>react-native</a></li>
            <li><a>alm</a></li>
            <li><a>beacon-ml</a></li>
          </ul>
        </div>
        <div className="tm-main uk-section uk-section-default">
          <div className="uk-container uk-container-small uk-position-relative">
            <div>
              <h1 className="uk-h1 tm-heading-fragment">React Native</h1>
              <p className="uk-text-lead">Readme content...</p>
            </div>
          </div>
        </div>
      </div>
    );
  }
}

export default App;
