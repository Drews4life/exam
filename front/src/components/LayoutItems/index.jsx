import React from 'react'
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';
import usePostResource from '../../hooks/usePostResource'
import { useHistory } from 'react-router-dom';

const useStyles = makeStyles({
  card: {
    minWidth: 275,
  },
  bullet: {
    display: 'inline-block',
    margin: '0 2px',
    transform: 'scale(0.8)',
  },
  title: {
    fontSize: 14,
  },
  pos: {
    marginBottom: 12,
  },
});

const LayoutItems = ({ item, page, fetchAgain }) => {
    const classes = useStyles();
    const history = useHistory();

    const onUpdateClick = () => history.push(`/${page}/${item.id}`)

    const onDeleteClick = async () => {
        await usePostResource(`/${page}/${item.id}`, "DELETE")
        fetchAgain();
    }

    return (
        <Card className={classes.card} variant="outlined">
          <CardContent>
            <Typography className={classes.title} color="textSecondary" gutterBottom>
              {item.author || item.name}
            </Typography>
            <Typography variant="h5" component="h2">
                {item.biography || item.title}
            </Typography>
            <Typography className={classes.pos} color="textSecondary">
              {item.year || item.dateOfBirth}
            </Typography>
          </CardContent>
            <CardActions>
                <Button size="small" onClick={onUpdateClick}>Update</Button>
                <Button size="small" onClick={onDeleteClick}>Delete</Button>
            </CardActions>
        </Card>
      );
}

export default LayoutItems