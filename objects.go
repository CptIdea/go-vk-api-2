package vk

type Updates struct {
	Updates []struct {
		Type   string
		Object struct {
			Message Message
		}
	}
}

type Message struct {
	Text            string
	Peer_id         int
	Id              int
	From_id         int
	Date            int
	Random_id       int
	Important       bool
	Payload         string
	Reply_message   *Message
	Keyboard        Keyboard
	Ref, Ref_source string
	Action          struct {
		Type      string
		Member_id int
		Text      string
		Email     string
		Photo     struct {
			Photo_50, Photo_100, Photo_200 string
		}
	}
	Geo struct {
		Type        string
		Coordinates struct {
			Latitude, Longitude float32
		}
		Place Place
	}
	Conversation_message_id int
	Attachments             []Attachments
	Fwd_messages            []*Message
}

type User struct {
	// Main Fields
	Id               int
	First_name       string
	Last_name        string
	Deactivated      string
	Is_closed        bool
	Can_acces_closed bool
	// Extra Fields
	About                     string
	Activities                string
	Bdate                     string
	Blacklisted               int
	Blacklisted_by_me         int
	Books                     string
	Can_post                  int
	Can_see_all_posts         int
	Can_see_audio             int
	Can_send_friend_request   int
	Can_write_private_message int
	Career                    struct {
		Group_id   int
		Company    string
		Country_id int
		City_id    int
		City_name  string
		From       int
		Until      int
		Position   string
	}
	City struct {
		Id    int
		Title string
	}
	Common_count int
	Connections  struct {
		Instagram   string
		Skype       string
		Facebook    string
		Twitter     string
		Livejournal string
	}
	Contacts struct {
		Mobile_phone string
		Home_phone   string
	}
	Counters struct {
		Albums, Videos, Audio, Photos, Notes, Friends, Groups, Online_friends, Mutual_friends, User_videos, Followers, Pages int
	}
	Country struct {
		Id    int
		Title string
	}
	Crop_photo struct {
		Photo struct {
			Id, Album_id, Owner_id, User_id, Date, Width, Height                    int
			Text, Photo_75, Photo_130, Photo_604, Photo_807, Photo_1280, Photo_2560 string
			Sizes                                                                   []struct {
				Url, Type     string
				Width, Height int
			}
		}
	}
	Domain    string
	Education struct {
		University, Faculty, Graduation int
		University_name, Faculty_name   string
	}
	Followers_count                                                                          int
	Friend_status                                                                            int
	Games                                                                                    string
	Has_mobile                                                                               int
	Has_photo                                                                                int
	Home_town                                                                                string
	Interests                                                                                string
	Is_favorite, Is_friend, Is_hidden_from_feed                                              int
	Last_name_nom, Last_name_gen, Last_name_dat, Last_name_acc, Last_name_ins, Last_name_abl string
	Last_seen                                                                                struct {
		Time     int
		Platform int
	}
	Lists       string
	Maiden_name string
	Military    struct {
		Unit                             string
		Unit_id, Country_id, From, Until int
	}
	Movies     string
	Music      string
	Nickname   string
	Occupation struct {
		Type string
		Id   int
		Name string
	}
	Online        int
	Online_mobile int
	Online_app    int
	Personal      struct {
		Political                                int
		Langs                                    []int
		Religion                                 string
		Inspired_by                              string
		People_main, Life_main, Smoking, Alcohol int
	}
	Photo_50, Photo_100, Photo_200_orig, Photo_200, Photo_400_orig, Photo_id, Photo_max, Photo_max_orig string
	Quotes                                                                                              string
	Relatives                                                                                           []struct {
		Id   int
		Name string
		Type string //Child Sibling Parent GrandParent GrandChild
	}
	Relation int
	Schools  []struct {
		Id, Country, City  int
		Name               string
		Year_from, Year_to int
		Year_Graduated     int
		Class              string
		Speciality         string
		Type               int    // Тип школы
		Type_str           string //0 — "школа"; 1 — "гимназия"; 2 —"лицей"; 3 — "школа-интернат"; 4 — "школа вечерняя"; 5 — "школа музыкальная"; 6 — "школа спортивная"; 7 — "школа художественная"; 8 — "колледж"; 9 — "профессиональный лицей"; 10 — "техникум"; 11 — "ПТУ"; 12 — "училище"; 13 — "школа искусств".
	}
	Screen_name  string
	Sex          int
	Site         string
	Timezone     int
	Trending     int
	Tv           string
	Universities []struct {
		Id, Country, City int
		Name              string
		Faculty           int
		Faculty_name      string
		Chair             int
		Chair_name        string
		Graduation        int
		Education_form    string
		Education_status  string
	}
	Verified     int
	Wall_default string
}
type Group struct {
	// Main
	Id                             int
	Name                           string
	Screen_name                    string
	Is_closed                      int
	Deactivated                    int
	Is_admin                       int
	Admin_level                    int
	Is_member                      int
	Is_advertiser                  int
	Invited_by                     int
	Type                           string
	Photo_50, Photo_100, Photo_200 string
	// Extra
	Activity                string
	Age_limits              int
	Can_create_topic        int
	Can_message             int
	Can_post                int
	Can_see_all_posts       int
	Can_upload_doc          int
	Can_upload_video        int
	Description             string
	Fixed_post              int
	Has_photo               int
	Is_favorite             int
	Is_hidden_from_feed     int
	Is_messages_blocked     int
	Main_album_id           int
	Main_section            int
	Member_status           int
	Members_count           int
	Place                   Place
	Public_date_label       string
	Site                    string
	Start_date, Finish_date int
	Status                  string
	Trending                int
	Verified                int
	Wall                    int
	Wiki_page               string
	Adresses                struct {
		Is_enabled     bool
		Main_adress_id int
	}
	Ban_info struct {
		End_date int
		Comment  string
	}
	City struct {
		Id    int
		Title string
	}
	Contacts []struct {
		User_id int
		Desc    string
		Phone   string
		Email   string
	}
	Counters struct {
		Photos,
		Albums,
		Audios,
		Videos,
		Topics,
		Docs int
	}
	Country struct {
		Id    int
		Title string
	}
	Cover struct {
		Enabled int
		Images  []struct {
			Url    string
			Width  int
			Height int
		}
	}
	Crop_photo struct {
		Photo Photo
		Crop  struct {
			X, Y, X1, Y1 float32
		}
		Rect struct {
			X, Y, X1, Y1 float32
		}
	}
	Links []struct {
		Id        int
		Url       string
		Name      string
		Desc      string
		Photo_50  string
		Photo_100 string
	}
	Market struct {
		Enabled       int
		Price_min     int
		Price_max     int
		Main_album_id int
		Contact_id    int
		Currency      Currency
		Currency_text string
	}
}
type Conversation struct {
	In_read, Out_read int
	Unread_count      int
	Important         bool
	Unanswered        bool
	Peer              struct {
		Id       int
		Type     string
		Local_id int
	}
	Push_settings struct {
		Disabled_untill  int
		Disabled_forever bool
		No_sound         bool
	}
	Can_write struct {
		Allowed bool
		Reason  int
	}
	Chat_settings struct {
		Members_count  int
		Title          string
		Pinned_message Message
		State          string
		Photo          struct {
			Photo_50, Photo_100, Photo_200 string
		}
		Active_ids       []int
		Is_group_channel bool
	}
}
type Chat struct {
	Id                             int
	Type                           string
	Admin_id                       int
	Users                          []int
	Photo_50, Photo_100, Photo_200 string
	Left                           int
	Kicked                         int
	Push_settings                  struct {
		Sound          int
		Disabled_until int
	}
}

type Attachments struct {
	Type          string
	Photo         Photo
	Audio         Audio
	Video         Video
	Sticker       Sticker
	Wall          Wall
	Note          Note
	Market        Market
	Market_album  Market_album
	Page          Page
	App           App
	Pool          Pool
	WallComment   WallComment
	Graffiti      Graffiti
	Audio_message Audio_message
	Document      Document
}
type Photo struct {
	Id            int
	Owner_Id      int
	Date          int
	User_id       int
	Album_id      int
	Sizes         []PhotoSizes
	Width, Height int
}
type PhotoSizes struct {
	Type, Url     string
	Width, Height int
}
type Audio struct {
	Artist    string
	Title     string
	Album_id  int
	Id        int
	Owner_id  int
	Duration  int
	Url       string
	Lyrics_id int
	Genre_id  int
	Date      int
	No_search int
	Is_hq     int
}
type Video struct {
	Id               int
	Owner_id         int
	Title            string
	Description      string
	Duration         int
	Photo_130        string
	Photo_320        string
	Photo_640        string
	Photo_800        string
	Photo_1280       string
	First_frame_130  string
	First_frame_320  string
	First_frame_640  string
	First_frame_800  string
	First_frame_1280 string
	Date             int
	Adding_date      int
	Views            int
	Comments         int
	Player           string
	Platform         string
	Can_edit         int
	Can_add          int
	Is_private       int
	Access_key       string
	Processing       int
	Live             int
	Upcoming         int
	Is_favorite      bool
}
type Document struct {
	Id       int
	Owner_id int
	Title    string
	Size     int
	Ext      string
	Url      string
	Date     int
	Type     int
	Preview  struct {
		Graffiti      Graffiti
		Audio_message Audio_message
		Photo         struct {
			Sizes PhotoSizes
		}
	}
}
type Graffiti struct {
	Src    string
	Width  int
	Height int
}
type Audio_message struct {
	Duration int
	Waveform []int
	Link_ogg string
	Link_mp3 string
}
type Link struct {
	Url         string
	Title       string
	Caption     string
	Description string
	Photo       Photo
	Product     Product
	
}
type Product struct {
	Price struct {
		Amount   int
		Currency Currency
		Text     string
	}
}
type Currency struct {
	Id   int
	Name string
}
type Sticker struct {
	Product_Id int
	Sticker_Id int
	Images     []struct {
		Url           string
		Width, Height int
	}
	Images_with_background []struct {
		Url           string
		Width, Height int
	}
}
type Wall struct {
	Id             int
	Owner_id       int
	From_id        int
	Created_by     int
	Date           int
	Text           string
	Reply_owner_id int
	Reply_post_id  int
	Friends_only   int
	Comments       struct {
		Count           int
		Can_post        int
		Groups_can_post int
		Can_close       bool
		Can_open        bool
	}
	Copyright string
	Likes     struct {
		Count       int
		User_likes  int
		Can_like    int
		Can_publish int
	}
	Reposts struct {
		Count         int
		User_reposted int
	}
	Views struct {
		Count int
	}
	Post_Type   string
	Post_source struct {
		Type     string
		Platform string
		Data     string
		Url      string
	}
	Attachments []Attachments
	Geo         struct {
		Type        string
		Coordinates string
		Place       Place
	}
	Signer_id     int
	Can_pin       int
	Can_delete    int
	Can_edit      int
	Is_pinned     int
	Marked_as_ads int
	Is_favorite   bool
	Postponed_id  int
}
type WallComment struct {
	Id               int
	From_id          int
	Date             int
	Text             string
	Reply_to_user    int
	Reply_to_comment int
	Attachments      []Attachments
	Parrents_stack   []int
	Thread           struct {
		Count             int
		Items             []*WallComment
		Can_post          bool
		Show_reply_button bool
		Groups_can_post   bool
	}
}
type Address struct {
	Id                 int
	Country_id         int
	City_id            int
	Title              string
	Address            string
	Additional_address string
	Latitude           float32
	Longitude          float32
	Metro_station_id   int
	Work_info_status   string
	Timetable          struct {
		Fri GroupTimetable
		Mon GroupTimetable
		Sat GroupTimetable
		Thu GroupTimetable
		Tue GroupTimetable
		Wed GroupTimetable
	}
}
type GroupTimetable struct {
	Close_time       int
	Open_time        int
	Break_close_time int
	Break_open_time  int
}
type Note struct {
	Id            int
	Owner_id      int
	Title         string
	Text          string
	Date          int
	Comments      int
	Read_comments int
	View_url      int
}
type Page struct {
	Id                           int
	Group_id                     int
	Creator_id                   int
	Title                        string
	Current_user_can_edit        int
	Current_user_can_edit_access int
	Who_can_view                 int
	Who_can_edit                 int
	Edited                       int
	Created                      int
	Editor_id                    int
	Views                        int
	Parent                       string
	Parent2                      string
	Source                       string
	Html                         string
	View_url                     string
}
type Market struct {
	// Main
	Id           int
	Owner_id     int
	Title        string
	Description  string
	Thumb_photo  string
	Date         int
	Availability int
	Is_favorite  bool
	Price        struct {
		Amount     int
		Old_amount int
		Text       string
		Currency   Currency
	}
	Category struct {
		Id      int
		Name    string
		Section struct {
			Id   int
			Name string
		}
	}
	// Extra
	Photos       []Photo
	Can_comment  int
	Can_repost   int
	Url          string
	Button_title string
	Likes        struct {
		User_likes int
		Count      int
	}
}
type Market_album struct {
	Id           int
	Owner_id     int
	Title        string
	Photo        Photo
	Count        int
	Updated_time int
}
type Order struct {
	Id                   int
	Group_id             int
	User_id              int
	Date                 int
	Variants_grouping_id int
	Is_main_variant      bool
	Cart_quantity        int
	Status               int
	Items_count          int
	Display_order_id     string
	Track_number         string
	Track_link           string
	Comment              string
	Adress               string
	Preview_order_items  []Market
	Property_values      []struct {
		Variant_id    int
		Variant_name  string
		Property_name string
	}
	Total_price struct {
		Amount   string
		Text     string
		Currency Currency
	}
}
type Board struct {
	Id            int
	Title         string
	Created       int
	Created_by    int
	Updated       int
	Is_closed     int
	Is_fixed      int
	Comments      int
	First_comment string
	Last_comment  string
}
type BoardComment struct {
	Id          int
	From_id     int
	Date        int
	Text        string
	Attachments []Attachments
	Likes       struct {
		Count      int
		User_likes int
		Can_like   int
	}
}
type App struct {
	// Main
	Id                       int
	Title                    string
	Icon_278                 string
	Icon_139                 string
	Icon_150                 string
	Icon_75                  string
	Banner_560               string
	Banner_1120              string
	Type                     string
	Section                  string
	Author_url               string
	Author_id                string
	Author_group             string
	Members_count            string
	Published_date           int
	Catalog_position         int
	International            int
	Leaderboard_type         int
	Genre_id                 int
	Genre                    string
	Platform_id              string
	Is_in_catalog            int
	Friends                  []int
	Installed                int
	Is_html5_app             int
	Screen_orientation       int
	Mobile_conrols_type      int
	Mobile_view_support_type int
	// Extra
	Description  string
	Screen_name  string
	Icon_16      string
	Screenshots  []Photo
	Push_enabled int
}
type Pool struct {
	Id         int
	Owner_id   int
	Created    int
	Question   string
	Votes      int
	Anonymous  bool
	Multiple   bool
	Answer_ids []int
	End_date   int
	Closed     bool
	Is_board   bool
	Can_edit   bool
	Can_report bool
	Can_share  bool
	Author_id  int
	Photo      Photo
	Friends    []int
	Background struct {
		Id     int
		Type   string
		Angle  int
		Color  string
		Width  int
		Height int
		Images []Photo
		Points []struct {
			Position float32
			Color    string
		}
	}
	Answers []struct {
		Id    int
		Text  string
		Votes int
		Rate  float32
	}
}
type Stats struct {
	Period_from int
	Period_to   int
	Visitors    struct {
		Views    int
		Visitors int
	}
	Reach struct {
		Reach             int
		Reach_subscribers int
		Mobile_reach      int
		Sex               []struct {
			Value string
			Count int
		}
		Age []struct {
			Value string
			Count int
		}
		Sex_age []struct {
			Value string
			Count int
		}
		Cities []struct {
			Name    string
			City_id interface{}
			Count   int
		}
		Countries []struct {
			Name       string
			Code       string
			Country_id int
			Count      int
		}
	}
}

type Place struct {
	Id                  int
	Title               string
	Latitude, Longitude float32
	Created             int
	Icon                string
	Country             string
	City                string
}
type Keyboard struct {
	One_time, Inline bool
	Buttons          []Button
}
type Button struct {
	Action []struct {
		Type     string
		Label    string
		Payload  string
		Link     string
		Hash     string
		App_id   int
		Owner_id int
	}
	Color string
}
