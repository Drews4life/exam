import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter, Route } from 'react-router-dom';
import { hot } from 'react-hot-loader/root';

import AuthorsPage from './pages/Authors'
import AuthorsUpdatePage from './pages/UpdateAuthor'
import BooskPage from './pages/Books'
import BooksUpdatePage from './pages/UpdateBook'
import MainPage from './pages/Main'
import SearchPage from './pages/Search'
import AddNewPage from './pages/AddNew'

import './index.scss'

const App = () => (
    <BrowserRouter>
        <Route path="/" component={MainPage} />
        <Route exact path="/authors" component={AuthorsPage} />
        <Route exact path="/authors/:id" component={AuthorsUpdatePage} />
        <Route exact path="/books" component={BooskPage} />
        <Route exact path="/books/:id" component={BooksUpdatePage} />
        <Route exact path="/search" component={SearchPage} />
        <Route exact path="/add/:type" component={AddNewPage} />
    </BrowserRouter>
)

const HotLoadedApp = hot(App)

ReactDOM.render(<HotLoadedApp/>, document.getElementById('app'));
