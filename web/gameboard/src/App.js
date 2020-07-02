import React from 'react';
import {
    HashRouter as Router,
    Switch,
    Route
} from "react-router-dom";
import Index from "./components/pages/index"
import Admin from "./components/pages/admin/index"

function App() {
    return (
        <div className="App">
            <Router>
                <Switch>
                    <Route path="/admin">
                        <Admin />
                    </Route>                
                    <Route exact path="/">
                        <Index />
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
