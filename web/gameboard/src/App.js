import React from 'react';
import {
    HashRouter as Router,
    Switch,
    Route,
    Link
} from "react-router-dom";
import Index from "./components/pages/index"

function App() {
    return (
        <div className="App">
            <Router>
                <Switch>
                    <Route path="/something">
                        <h2>hello</h2>
                    </Route>                
                    <Route exact path="/">
                        <Index component={Route} />
                    </Route>
                    <Route path="/">
                        <h2>Sad</h2>
                    </Route>
                </Switch>
            </Router>
        </div>
    );
}

export default App;
