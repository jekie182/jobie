package interfaces


//IResult type of interface has include each object orcollection of objects which we are returning as response
type IResult interface{
	var Succes bool;
	var []string Errors;
	var IEntity Entity;
}