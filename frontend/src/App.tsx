import React from 'react';
import {LandingPage} from "./containers/LandingPage";
import {Route, Switch} from "react-router-dom";
import Login from "./containers/Login";
import 'bootstrap/dist/css/bootstrap.min.css';
import SignUp from "./containers/Signup";

/*
Notes: Conditional rendering for the "/" route
When the user is logged in, it should automatically bring them to the main page rather than the landing one
 */
function App() {
    return (
        <div className="App">
            <Switch>
                <Route exact path={"/"} component={LandingPage}/>
                <Route exact path={"/login"} component={Login}/>
                <Route exact path={"/signup"} component={SignUp}/>
            </Switch>
        </div>
    );
}

export default App;
