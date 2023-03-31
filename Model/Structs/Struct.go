package Structs

type Post struct {
	Id_post   int `json:"Id_post"`
	Name_Post int `json:"Name_Post"`
}
type Marks struct {
	ID_Marks   int `json:"ID_Marks"`
	Name_Marks int `json:"Name_Marks"`
}
type Configurat struct {
	ID_Configurat   int `json:"ID_Configurat"`
	Name_Configurat int `json:"Name_Configurat"`
	Price           int `json:"Price"`
}
type Model struct {
	ID_Model      int `json:"ID_Model"`
	Name_Model    int `json:"Name_Model"`
	Marks_ID      int `json:"Marks_ID"`
	Configurat_ID int `json:"Configurat_ID"`
}
type Gears struct {
	ID_Gears   int `json:"ID_Gears"`
	Name_Gears int `json:"Name_Gears"`
	Price      int `json:"Price"`
	Model_ID   int `json:"Model_ID"`
}
type Roles struct {
	ID_Roles int `json:"ID_Roles"`
	NameR    int `json:"NameR"`
}

type Users struct {
	ID_Users int    `json:"ID_Users"`
	Login    string `json:"Login"`
	Passwor  string `json:"Passwor"`
	Roles_ID int    `json:"Roles_ID"`
}
type Cont_info struct {
	ID_Cont_info int `json:"ID_Cont_info"`
	Title        int `json:"Title"`
	Adress       int `json:"Adress"`
	Numbe        int `json:"Numbe"`
	Desribe      int `json:"Desribe"`
}

type Employee struct {
	ID_Employee     int `json:"ID_Employee"`
	first_Name      int `json:"first_Name"`
	Name_Employee   int `json:"Name_Employee"`
	Middle_Name     int `json:"Middle_Name"`
	Date_Bithdais   int `json:"Date_Bithdais"`
	Employee_Number int `json:"Employee_Number"`
	Employee_Email  int `json:"Employee_Email"`
	Post_ID         int `json:"Post_ID"`
	Users_ID        int `json:"Users_ID"`
	Active          int `json:"Active"`
}
type Soiscatel struct {
	ID_Soiscatel  int `json:"ID_Soiscatel"`
	first_Name    int `json:"first_Name"`
	Name_s        int `json:"Name_s"`
	Middle_Name_s int `json:"Middle_Name_s"`
	Confirm       int `json:"Confirm"`
	Roles_ID      int `json:"Roles_ID"`
	User_ID       int `json:"User_ID"`
}

type Client struct {
	ID_Client     int `json:"ID_Client"`
	first_Name    int `json:"first_Name"`
	Middle_Name_s int `json:"Middle_Name_s"`
	Post          int `json:"Post"`
	Date_Bithdais int `json:"Date_Bithdais"`
	Client_Number int `json:"Client_Number"`
	Client_Email  int `json:"Client_Email"`
	User_ID       int `json:"User_ID"`
	Active        int `json:"Active"`
}

type Cheak struct {
	ID_Cheak    int `json:"ID_Cheak"`
	Numbers     int `json:"Numbers"`
	Dates       int `json:"Dates"`
	Counts      int `json:"Counts"`
	Employee_ID int `json:"Employee_ID"`
	Client_ID   int `json:"Client_ID"`
	Model_ID    int `json:"Model_ID"`
}
