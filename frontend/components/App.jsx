import React from "react"
import PropTypes from 'prop-types'
import axios from 'axios'
import renderHTML from 'react-render-html'
import Pagination from "react-js-pagination"


const SERVER_URL = 'http://' + process.env.API_URL + '/api/v1'
console.log(SERVER_URL)

class App extends React.Component {

    static propTypes = {
        posts: PropTypes.array.isRequired
    }

    constructor(props) {
      super(props);
      this.state = {data: [], activePage: 1, count: 1}
      this.handlePageChange = this.handlePageChange.bind(this)
    }

    componentWillMount() {
        axios.get(SERVER_URL + "/")
          .then(response => this.setState({ data: response.data.Posts, count: response.data.Count}))
    }
  
    handlePageChange(pageNumber) {
      this.setState({ activePage: pageNumber });
      axios.get(`${SERVER_URL}/?page=${pageNumber}`)
        .then(response => this.setState({ data: response.data.Posts }))
      window.scrollTo(0, 0)
    }

    handleGetPost(id){
      axios.get(`${SERVER_URL}/post/${id}`)
        .then(response => this.setState({ data: response.data.Posts }))
    }

    render() {
      console.log(this.state.data)
      const posts = this.state.data.map(item => {
        return (
          <div class="row">
            <div class="panel panel-default">
              <div class="panel-body">
                <h2 onClick={() => this.handleGetPost(item.Id)}><a href="#">{item.Title}</a></h2>
                <div id="post_body">{renderHTML(item.Text)}</div>
              </div>
            </div>
          </div>
          )
        }
      )
      return (
        <div className="component-app">
            {posts}
            <div>
              <Pagination
                activePage={this.state.activePage}
                itemsCountPerPage={5}
                totalItemsCount={this.state.count}
                pageRangeDisplayed={5}
                onChange={this.handlePageChange}
              />
            </div>
        </div>
      );
    }
  }
  export default App;