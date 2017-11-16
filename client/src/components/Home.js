import React, { Component } from 'react'
import RepoCard from './RepoCard'
import { Row } from '../theme/globalStyle'

const repos = [
  {
    name: 'repo1',
    username: 'wayanjimmy',
    description: 'Karepmu',
    url: 'https://github.com/wayanjimmy/crazy-weekend',
  },
  {
    name: 'repo1',
    username: 'wayanjimmy',
    description: 'Karepmu',
    url: 'https://github.com/wayanjimmy/crazy-weekend',
  },
  {
    name: 'repo1',
    username: 'wayanjimmy',
    description: 'Karepmu',
    url: 'https://github.com/wayanjimmy/crazy-weekend',
  },
  {
    name: 'repo1',
    username: 'wayanjimmy',
    description: 'Karepmu',
    url: 'https://github.com/wayanjimmy/crazy-weekend',
  },
]

class Home extends Component {
  render() {
    return (
      <div>
        <Row>
          <div className="col">
            <form className="form-inline mt-2 mt-md-0">
              <input
                className="form-control mr-sm-2"
                type="text"
                placeholder="Search by language"
                aria-label="Search"
              />
            </form>
          </div>
        </Row>
        <Row>{repos.map((repo, index) => <RepoCard key={index} repo={repo} />)}</Row>
      </div>
    )
  }
}

export default Home
