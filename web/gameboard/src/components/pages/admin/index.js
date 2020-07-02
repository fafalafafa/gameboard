import React, { Component } from "react";
import { GetCharacters } from "../../middlewares/admin"
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import config from "../../../config/role-maps.json";

import "./indes.css";


class Index extends Component {

    constructor(props) {
        super(props)
        this.state = {
            name: "",
            isLoaded: false,
            characters: []
        }

    }

    componentWillMount() {
        GetCharacters().then(({ data }) => {
            this.setState({
                characters: data.characters
            })
        })
    }

    render() {

        const { characters } = this.state;
        return (
            <TableContainer component={Paper} className={"admin-table"}>
                <Table size="small" aria-label="a dense table">
                    <TableHead>
                        <TableRow>
                            <TableCell>Player</TableCell>
                            <TableCell align="right">Image</TableCell>
                            <TableCell align="right">Name</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {characters.map((row) => (
                            <TableRow key={row.name}>
                                <TableCell component="th" scope="row">
                                    {row.Name}
                                </TableCell>
                                <TableCell align="right">
                                    <img className="character-portrait" src={`/characters/${row.Character}.jpg`} />
                                </TableCell>
                                <TableCell align="right">{config[row.Character].text}</TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </TableContainer>
        );
    }

}

export default Index;