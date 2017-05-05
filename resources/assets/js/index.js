import React from 'react';
import { render } from 'react-dom';
import { Provider } from 'react-redux';
import { applyMiddleware, createStore } from 'redux';
import logger from 'redux-logger';
import klengApp from './reducers';
import App from './containers/App';

window.$ = window.jQuery = require('jquery');

const store = createStore(
  klengApp,
  applyMiddleware(logger)
);

render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('app')
);
