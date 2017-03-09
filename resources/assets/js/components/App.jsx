import React, { Component } from 'react';
import Nav from './Nav';

class App extends Component {
  render() {
    return (
      <div>
        <Nav />

        <div className="container">

          <div className="row">
            <div className="col-lg-12 text-center">
              <h1>A Bootstrap Starter Template</h1>
              <p className="lead">
                Hello World!
              </p>
              <ul className="list-unstyled">
                <li>Bootstrap v3.3.7</li>
                <li>jQuery v1.11.1</li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    );
  }
}

export default App;
