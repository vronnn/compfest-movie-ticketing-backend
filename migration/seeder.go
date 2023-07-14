package migration

import (
	"errors"
	"fmt"
	"time"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	if err := ListBankSeeder(db); err != nil {
		return err
	}

	if err := ListUserSeeder(db); err != nil {
		return err
	}

	if err := ListMovieSeeder(db); err != nil {
		return err
	}

	if err := ListTimeMovieSeeder(db); err != nil {
		return err
	}

	if err := ListStudioSeeder(db); err != nil {
		return err
	}

	return nil
}

func ListBankSeeder(db *gorm.DB) error {
	var listBank = []entities.ListBank{
		{
			ID:   1,
			Name: "BCA",
		},
		{
			ID:   2,
			Name: "BNI",
		},
		{
			ID:   3,
			Name: "BRI",
		},
		{
			ID:   4,
			Name: "Mandiri",
		},
		{
			ID:   5,
			Name: "OVO",
		},
		{
			ID:   6,
			Name: "Gopay",
		},
	}

	hasTable := db.Migrator().HasTable(&entities.ListBank{})
	if !hasTable {
		if err := db.AutoMigrate(&entities.ListBank{}); err != nil {
			return err
		}
	}

	for _, data := range listBank {
		var bank entities.ListBank
		err := db.Where(&entities.ListBank{ID: data.ID}).First(&bank).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

func ListTimeMovieSeeder(db *gorm.DB) error {
	var listTimeMovie = []entities.TimeMovie{
		{
			ID:   1,
			Time: "12.05",
			Type: 1,
		},
		{
			ID:   2,
			Time: "14.30",
			Type: 1,
		},
		{
			ID:   3,
			Time: "16.05",
			Type: 1,
		},
		{
			ID:   4,
			Time: "18.30",
			Type: 1,
		},
		{
			ID:   5,
			Time: "20.05",
			Type: 1,
		},
		{
			ID:   6,
			Time: "22.30",
			Type: 1,
		},
		{
			ID:   7,
			Time: "10.05",
			Type: 2,
		},
		{
			ID:   8,
			Time: "12.30",
			Type: 2,
		},
		{
			ID:   9,
			Time: "14.05",
			Type: 2,
		},
		{
			ID:   10,
			Time: "16.30",
			Type: 2,
		},
		{
			ID:   11,
			Time: "18.05",
			Type: 2,
		},
		{
			ID:   12,
			Time: "20.30",
			Type: 2,
		},
		{
			ID:   13,
			Time: "22.05",
			Type: 2,
		},
		{
			ID:   14,
			Time: "09.05",
			Type: 3,
		},
		{
			ID:   15,
			Time: "11.30",
			Type: 3,
		},
		{
			ID:   16,
			Time: "13.05",
			Type: 3,
		},
		{
			ID:   17,
			Time: "15.30",
			Type: 3,
		},
		{
			ID:   18,
			Time: "17.05",
			Type: 3,
		},
		{
			ID:   19,
			Time: "19.30",
			Type: 3,
		},
		{
			ID:   20,
			Time: "11.05",
			Type: 4,
		},
		{
			ID:   21,
			Time: "13.30",
			Type: 4,
		},
		{
			ID:   22,
			Time: "15.05",
			Type: 4,
		},
		{
			ID:   23,
			Time: "17.30",
			Type: 4,
		},
		{
			ID:   24,
			Time: "19.05",
			Type: 4,
		},
		{
			ID:   25,
			Time: "21.30",
			Type: 4,
		},
		{
			ID:   26,
			Time: "10.05",
			Type: 5,
		},
		{
			ID:   27,
			Time: "12.30",
			Type: 5,
		},
		{
			ID:   28,
			Time: "14.05",
			Type: 5,
		},
		{
			ID:   29,
			Time: "16.30",
			Type: 5,
		},
		{
			ID:   30,
			Time: "18.05",
			Type: 5,
		},
	}

	for _, data := range listTimeMovie {
		var timeMovie entities.TimeMovie
		err := db.Where(&entities.TimeMovie{ID: data.ID}).First(&timeMovie).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

func ListStudioSeeder(db *gorm.DB) error {
	var listPlace = []entities.Place{
		{
			ID:   1,
			Name: "CGV",
		},
		{
			ID:   2,
			Name: "XXI",
		},
		{
			ID:   3,
			Name: "Cinepolis",
		},
		{
			ID:   4,
			Name: "Cinemaxx",
		},
		{
			ID:   5,
			Name: "New Star Cineplex",
		},
		{
			ID:   6,
			Name: "Flix Cinema",
		},
		{
			ID:   7,
			Name: "Hiflix",
		},
		{
			ID:   8,
			Name: "Platinum Cineplex",
		},
	}

	for _, data := range listPlace {
		var place entities.Place
		err := db.Where(&entities.Place{ID: data.ID}).First(&place).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

func ListUserSeeder(db *gorm.DB) error {
	var listUser = []entities.User{
		{
			Nama:         "Admin",
			NoTelp:       "081234567890",
			Email:        "Admin@gmail.com",
			Password:     "Admin123",
			Age:          20,
			TanggalLahir: "08/15/2003",
			Role:         "Admin",
			Saldo:        0,
		},
		{
			Nama:         "User",
			NoTelp:       "081234567890",
			Email:        "User@gmail.com",
			Password:     "User123",
			Age:          20,
			TanggalLahir: "08/15/2003",
			Role:         "User",
			Saldo:        0,
		},
	}

	hasTable := db.Migrator().HasTable(&entities.User{})
	if !hasTable {
		if err := db.AutoMigrate(&entities.User{}); err != nil {
			return err
		}
	}

	for _, data := range listUser {
		var user entities.User
		err := db.Where(&entities.User{Email: data.Email}).First(&user).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		isData := db.Find(&user, "email = ?", data.Email).RowsAffected
		if isData == 0 {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

func ListMovieSeeder(db *gorm.DB) error {
	var listMovie = []dto.MovieCreateDTO{
		{
			Title:       "Fast X",
			Description: "Dom Toretto dan keluarganya menjadi sasaran putra gembong narkoba Hernan Reyes yang pendendam.",
			ReleaseDate: "2023-05-17",
			PosterUrl:   "https://image.tmdb.org/t/p/w500/fiVW06jE7z9YnO4trhaMEdclSiC.jpg",
			AgeRating:   15,
			TicketPrice: 53000,
		},
		{
			Title:       "John Wick: Chapter 4",
			Description: "ohn Wick mengungkap jalan untuk mengalahkan The High Table. Tapi sebelum dia bisa mendapatkan kebebasannya, Wick harus berhadapan dengan musuh baru dengan aliansi kuat di seluruh dunia dan kekuatan yang mengubah teman lama menjadi musuh.",
			ReleaseDate: "2023-03-22",
			PosterUrl:   "https://image.tmdb.org/t/p/w500/vZloFAK7NmvMGKE7VkF5UHaz0I.jpg",
			AgeRating:   10,
			TicketPrice: 60000,
		},
		{
			Title:       "The Super Mario Bros. Movie",
			Description: "Ketika sedang bekerja di bawah tanah untuk memperbaiki pipa air, Mario dan Luigi, yang merupakan tukang ledeng dari Brooklyn, tiba-tiba terhisap ke dalam pipa misterius dan masuk ke dunia yang sangat berbeda. Mereka berada di tempat yang ajaib dan aneh. Tapi sayangnya, mereka terpisah satu sama lain. Mario memulai petualangan besar untuk mencari dan menemukan Luigi.",
			ReleaseDate: "2023-04-05",
			PosterUrl:   "https://image.tmdb.org/t/p/w500/qNBAXBIQlnOThrVvA6mA2B5ggV6.jpg",
			AgeRating:   14,
			TicketPrice: 49000,
		},
		{
			Title:       "Avatar: The Way of Water",
			Description: "Jake Sully tinggal bersama keluarga barunya di planet Pandora. Setelah ancaman kembali datang, Jake harus bekerja dengan Neytiri dan pasukan ras Na'vi untuk melindungi planet mereka.",
			ReleaseDate: "2022-12-14",
			PosterUrl:   "https://image.tmdb.org/t/p/w500/t6HIqrRAclMCA60NsSmeqe9RmNV.jpg",
			AgeRating:   12,
			TicketPrice: 53000,
		},
		{
			Title:       "Guardians of the Galaxy Vol. 3",
			Description: "Peter Quill masih trauma karena kehilangan Gamora. Ia perlu mengumpulkan timnya untuk melindungi alam semesta dan salah satu anggota mereka. Jika mereka gagal, Guardian akan berakhir.",
			ReleaseDate: "2023-05-03",
			PosterUrl:   "https://image.tmdb.org/t/p/w500/nAbpLidFdbbi3efFQKMPQJkaZ1r.jpg",
			AgeRating:   12,
			TicketPrice: 41000,
		},
		{
			Title:       "Ant-Man and the Wasp: Quantumania",
			Description: "Scott Lang dan Hope van Dyne adalah pasangan pahlawan super. Mereka pergi bersama orang tua Hope, Janet van Dyne dan Hank Pym, serta anak perempuan Scott, Cassie Lang, untuk menjelajahi Alam Kuantum. Di sana, mereka bertemu dengan makhluk-makhluk aneh dan memulai petualangan yang tak terduga. Petualangan ini akan menguji batas-batas mereka.",
			ReleaseDate: "2023-02-15",
			PosterUrl:   "https://image.tmdb.org/t/p/w500/g0OWGM7HoIt866Lu7yKohYO31NU.jpg",
			AgeRating:   12,
			TicketPrice: 51000,
		},
		{
			Title:       "The Pope's Exorcist",
			Description: "Pastor Gabriele Amorth, yang memimpin tim pengusir setan di Vatikan, menginvestigasi kasus kekerasan roh jahat yang menghantui seorang anak laki-laki. Dalam penyelidikannya, ia secara tak terduga menemukan rahasia tua yang disembunyikan oleh Vatikan selama berabad-abad.",
			ReleaseDate: "2023-04-05",
			PosterUrl:   "https://image.tmdb.org/t/p/w500/gNPqcv1tAifbN7PRNgqpzY8sEJZ.jpg",
			AgeRating:   13,
			TicketPrice: 51000,
		},
		{
			Title:       "To Catch a Killer",
			Description: "Baltimore. Malam tahun baru. Seorang petugas polisi yang berbakat tetapi bermasalah (Shailene Woodley) direkrut oleh kepala penyelidik FBI (Ben Mendelsohn) untuk membantu membuat profil dan melacak individu yang terganggu yang meneror kota.",
			ReleaseDate: "2023-04-06",
			PosterUrl:   "https://image.tmdb.org/t/p/w500/mFp3l4lZg1NSEsyxKrdi0rNK8r1.jpg",
			AgeRating:   15,
			TicketPrice: 47000,
		},
		{
			Title:       "Transformers: Age of Extinction",
			Description: "Lima tahun setelah Chicago dihancurkan, manusia berbalik melawan robot. Namun seorang ayah tunggal dan penemu membangkitkan robot yang dapat menyelamatkan dunia.",
			ReleaseDate: "2014-06-25",
			PosterUrl:   "https://image.tmdb.org/t/p/w500/jyzrfx2WaeY60kYZpPYepSjGz4S.jpg",
			AgeRating:   11,
			TicketPrice: 54000,
		},
		{
			Title:       "Puss in Boots: The Last Wish",
			Description: "Puss in Boots menemukan fakta bahwa kecintaannya pada petualangan telah merenggut nyawanya: dia telah menghabiskan delapan dari sembilan nyawanya. Puss kini memulai petualangan epik untuk menemukan harapan terakhir untuk memulihkan sembilan nyawanya.",
			ReleaseDate: "2022-12-07",
			PosterUrl:   "https://image.tmdb.org/t/p/w500/kuf6dutpsT0vSVehic3EZIqkOBt.jpg",
			AgeRating:   11,
			TicketPrice: 51000,
		},
		{
			Title:       "Scream VI",
			Description: "Setelah pembunuhan terbaru oleh Ghostface, keempat orang yang selamat pergi dari Woodsboro dan memulai hidup baru.",
			ReleaseDate: "2023-03-08",
			PosterUrl:   "https://image.tmdb.org/t/p/w500/wDWwtvkRRlgTiUr6TyLSMX8FCuZ.jpg",
			AgeRating:   12,
			TicketPrice: 36000,
		},
		{
			Title:       "Black Adam",
			Description: "Hampir 5.000 tahun setelah dia dianugerahi kekuatan maha kuasa para dewa Mesirâ€”dan dipenjara dengan cepatâ€”Black Adam dibebaskan dari makam duniawinya, siap untuk melepaskan bentuk keadilannya yang unik di dunia modern.",
			ReleaseDate: "2022-10-19",
			PosterUrl:   "https://image.tmdb.org/t/p/w500/A5imhXiFF3AL9RRA4FBzNDFmfgW.jpg",
			AgeRating:   10,
			TicketPrice: 42000,
		},
		{
			Title:       "Dungeons & Dragons: Honor Among Thieves",
			Description: "Seorang pencuri menawan dan sekelompok petualang yang unik melakukan pencurian besar-besaran untuk mencuri relik yang hilang. Namun, segalanya menjadi kacau ketika mereka berjumpa dengan orang yang salah.",
			ReleaseDate: "2023-03-23",
			PosterUrl:   "https://image.tmdb.org/t/p/w500/A7AoNT06aRAc4SV89Dwxj3EYAgC.jpg",
			AgeRating:   12,
			TicketPrice: 38000,
		},
		{
			Title:       "Peter Pan & Wendy",
			Description: "Wendy Darling adalah seorang gadis kecil yang takut pergi dari rumah masa kecilnya. Suatu hari, dia bertemu dengan Peter Pan, seorang anak laki-laki yang tidak mau tumbuh dewasa. Mereka bersama saudara-saudaranya dan peri kecil bernama Tinker Bell pergi ke dunia ajaib yang disebut Neverland. Di sana, mereka menghadapi Kapten Hook, seorang bajak laut jahat, dan mengalami petualangan seru yang akan mengubah hidup Wendy selamanya.",
			ReleaseDate: "2023-04-20",
			PosterUrl:   "https://image.tmdb.org/t/p/w500/9NXAlFEE7WDssbXSMgdacsUD58Y.jpg",
			AgeRating:   13,
			TicketPrice: 35000,
		},
		{
			Title:       "Spider-Man: No Way Home",
			Description: "Peter Parker menghadapi masalah besar. Hal ini terjadi setelah identitasnya sebagai Spiderman terungkap. Dengan kepergian Tony Stark, Peter Parker pun harus meminta bantuan Doctor Strange agar semua orang bisa melupakan identitasnya sebagai manusia laba-laba.",
			ReleaseDate: "2021-12-15",
			PosterUrl:   "https://image.tmdb.org/t/p/w500/uJYYizSuA9Y3DCs0qS4qWvHfZg4.jpg",
			AgeRating:   15,
			TicketPrice: 56000,
		},
		{
			Title:       "Black Panther: Wakanda Forever",
			Description: "Rakyat Wakanda kali ini akan berjuang untuk melindungi negerinya dari campur tangan kekuatan dunia setelah kematian sang Raja T'Challa.",
			ReleaseDate: "2022-11-09",
			PosterUrl:   "https://image.tmdb.org/t/p/w500/sv1xJUazXeYqALzczSZ3O6nkH75.jpg",
			AgeRating:   13,
			TicketPrice: 39000,
		},
		{
			Title:       "Transformers: The Last Knight",
			Description: "Di tengah ketidakhadiran Optimus Prime, umat manusia berperang melawanTransformers untuk mempertahankan eksistensinya. Cade Yeager membentuk kerjasama dengan Bumblebee, seorang bangsawan Inggris dan seorang professor dari Oxford untuk mempelajari mengapa Transformers selalu kembali ke planet bumi.",
			ReleaseDate: "2017-06-16",
			PosterUrl:   "https://image.tmdb.org/t/p/w500/s5HQf2Gb3lIO2cRcFwNL9sn1o1o.jpg",
			AgeRating:   12,
			TicketPrice: 52000,
		},
		{
			Title:       "Renfield",
			Description: "Setelah bertahun-tahun sebagai hamba Dracula yang merasa jenuh dan lelah, Renfield menemukan harapan baru dalam hidupnya. Dia jatuh cinta pada Rebecca Quincy, seorang polisi lalu lintas yang energik dan sering marah. Kesempatan ini bisa menjadi penebusan baginya.",
			ReleaseDate: "2023-04-07",
			PosterUrl:   "https://image.tmdb.org/t/p/w500/2OaprROMZZeiWsydjGUIkXrv2Z3.jpg",
			AgeRating:   14,
			TicketPrice: 51000,
		},
		{
			Title:       "Cocaine Bear",
			Description: "Sekelompok polisi, penjahat, turis, dan remaja eksentrik berkumpul di hutan Georgia tempat beruang hitam besar mengamuk setelah menelan kokain secara tidak sengaja.",
			ReleaseDate: "2023-02-22",
			PosterUrl:   "https://image.tmdb.org/t/p/w500/gOnmaxHo0412UVr1QM5Nekv1xPi.jpg",
			AgeRating:   12,
			TicketPrice: 53000,
		},
		{
			Title:       "Prey",
			Description: "Di Comanche Nation pada tahun 1717, seorang pejuang yang ganas dan sangat terampil bernama Naru mengetahui bahwa mangsa yang dia intai adalah alien yang sangat berkembang dengan persenjataan berteknologi maju.",
			ReleaseDate: "2022-08-02",
			PosterUrl:   "https://image.tmdb.org/t/p/w500/ujr5pztc1oitbe7ViMUOilFaJ7s.jpg",
			AgeRating:   10,
			TicketPrice: 42000,
		},
		{
			Title:       "Fall",
			Description: "Untuk sahabat Becky dan Hunter, hidup adalah tentang menaklukkan ketakutan dan mendorong batas. Tetapi setelah mereka mendaki 2.000 kaki ke puncak menara radio terpencil yang ditinggalkan, mereka menemukan diri mereka terdampar tanpa jalan turun. Sekarang keterampilan panjat ahli Becky dan Hunter akan diuji saat mereka mati-matian berjuang untuk bertahan hidup dari unsur-unsur, kurangnya persediaan, dan ketinggian yang menyebabkan vertigo.",
			ReleaseDate: "2022-08-11",
			PosterUrl:   "https://image.tmdb.org/t/p/w500/v28T5F1IygM8vXWZIycfNEm3xcL.jpg",
			AgeRating:   11,
			TicketPrice: 39000,
		},
		{
			Title:       "Avatar",
			Description: "Pada abad ke-22, seorang Marinir lumpuh dikirim ke Pandora bulan pada misi yang unik, tetapi menjadi terpecah antara mengikuti perintah dan melindungi peradaban alien.",
			ReleaseDate: "2009-12-15",
			PosterUrl:   "https://image.tmdb.org/t/p/w500/kyeqWdyUXW608qlYkRqosgbbJyK.jpg",
			AgeRating:   13,
			TicketPrice: 37000,
		},
		{
			Title:       "Split",
			Description: "Ketika ketiga gadis remaja sedang menunggu ayah mereka di dalam mobil, seorang pria misterius menculik dan membawa mereka ke dalam sebuah bunker. Sang penculik yang bernama Kevin (James McAvoy) adalah seorang pria dengan gangguan jiwa yang membuatnya memiliki 23 kepribadian yang berbeda, yang diantaranya adalah seorang wanita dan anak berumur 9 tahun yang bernama Hedwig.  Sebagai salah satu gadis yang diculik, Casey berusaha meloloskan diri dengan meyakinkan salah satu kepribadian Kevin untuk melepaskan mereka. Akan tetapi hal tersebut tidaklah mudah, terlebih setelah Hedwig memperingatkan mereka akan the Beast yang merupakan kepribadian Kevin yang paling berbahaya.",
			ReleaseDate: "2017-01-19",
			PosterUrl:   "https://image.tmdb.org/t/p/w500/lli31lYTFpvxVBeFHWoe5PMfW5s.jpg",
			AgeRating:   10,
			TicketPrice: 45000,
		},
		{
			Title:       "Top Gun: Maverick",
			Description: "Setelah lebih dari tiga puluh tahun mengabdi sebagai salah satu penerbang top Angkatan Laut, dan menghindari kenaikan pangkat yang akan menjatuhkannya, Pete Maverick Mitchell mendapati dirinya melatih satu detasemen lulusan TOP GUN untuk misi khusus yang tidak ada kehidupan. pilot pernah melihat.",
			ReleaseDate: "2022-05-24",
			PosterUrl:   "https://image.tmdb.org/t/p/w500/jeGvNOVMs5QIU1VaoGvnd3gSv0G.jpg",
			AgeRating:   14,
			TicketPrice: 57000,
		},
		{
			Title:       "Thor: Love and Thunder",
			Description: "'Thor: Love and Thunder'menceritakan Thor (Chris Hemsworth) dalam sebuah perjalanan yang belum pernah ia jalani â€“ pencariankedamaian batin. Namun, masa pensiunnya terganggu oleh seorang pembunuh galaksi yang dikenal sebagai Gorr sang Dewa Jagal (Christian Bale), yang ingin memusnahkan para dewa. Untuk mengatasi ancaman, Thor meminta bantuan Raja Valkyrie (Tessa Thompson), Korg (Taika Waititi), dan mantan kekasihnya Jane Foster (Natalie Portman), yang secara mengejutkan dan misterius berhasil menggunakan palu ajaibnya, Mjolnir, sebagai Mighty Thor. Bersama, mereka memulai petualangan kosmik yang mendebarkan untuk mengungkap misteri pembalasan Dewa Jagal dan menghentikannya sebelum terlambat.",
			ReleaseDate: "2022-07-06",
			PosterUrl:   "https://image.tmdb.org/t/p/w500/pIkRyD18kl4FhoCNQuWxWu5cBLM.jpg",
			AgeRating:   12,
			TicketPrice: 35000,
		},
		{
			Title:       "Sonic the Hedgehog 2",
			Description: "Alur cerita film Sonic the Hedgehog 2 bermula ketika Sonic menetap di Green Hills. Ia memutuskan menetap di sana agar bisa merasakan lebih banyak kebebasan. Ditambah lagi, Tom dan Maddie setuju untuk meninggalakannya di rumah ketika mereka pergi untuk berlibur. Namun sayangnya, tidak lama setelah mereka pergi Dr. Robotnik sang musuh bubuyutan si landak biru itu kembali ke bumi. Kali ini Dr. Robotnik datang dengan pasukan baru, Knuckles. Tujuan mereka datang kembali adalah untuk mencari Master Emerald yang memiliki kekuatan super. Kekuatan super itu bisa membangun dan menghancurkan peradaban di bumi. Atas hal ini, Sonic pun mencari strategi agar bisa menggagalkan rencara Dr. Robotnik. Strategi yang dilakukan oleh Sonic ialah bekerjasama dengan sahabatnya, Tails. Kemudian bersama dengan Tails, Sonic memulai perjalanan untuk menemukan Master Emerald. Semua itu dilakukan dengan cepat, sebelum Master Emerald jatuh ke tangan yang salah.",
			ReleaseDate: "2022-04-08",
			PosterUrl:   "https://image.tmdb.org/t/p/w500/6DrHO1jr3qVrViUO6s6kFiAGM7.jpg",
			AgeRating:   12,
			TicketPrice: 45000,
		},
		{
			Title:       "Avengers: Infinity War",
			Description: "Karena Avengers dan sekutunya terus melindungi dunia dari ancaman yang terlalu besar untuk ditangani oleh seorang pahlawan, bahaya baru telah muncul dari bayangan kosmik: Thanos. Seorang lalim penghujatan intergalaksi, tujuannya adalah untuk mengumpulkan semua enam Batu Infinity, artefak kekuatan yang tak terbayangkan, dan menggunakannya untuk menimbulkan kehendak memutar pada semua realitas. Segala sesuatu yang telah diperjuangkan oleh Avengers telah berkembang hingga saat ini - nasib Bumi dan keberadaannya sendiri tidak pernah lebih pasti.",
			ReleaseDate: "2018-04-25",
			PosterUrl:   "https://image.tmdb.org/t/p/w500/7WsyChQLEftFiDOVTGkv3hFpyyt.jpg",
			AgeRating:   10,
			TicketPrice: 46000,
		},
		{
			Title:       "The Whale",
			Description: "Seorang guru bahasa Inggris yang tertutup dan gemuk mencoba untuk berhubungan kembali dengan putri remajanya yang terasing.",
			ReleaseDate: "2022-12-09",
			PosterUrl:   "https://image.tmdb.org/t/p/w500/jQ0gylJMxWSL490sy0RrPj1Lj7e.jpg",
			AgeRating:   15,
			TicketPrice: 55000,
		},
		{
			Title:       "The Batman",
			Description: "Ketika seorang pembunuh berantai sadis mulai membunuh tokoh-tokoh politik penting di Gotham, Batman terpaksa menyelidiki korupsi tersembunyi di kota itu dan mempertanyakan keterlibatan keluarganya.",
			ReleaseDate: "2022-03-01",
			PosterUrl:   "https://image.tmdb.org/t/p/w500/seyWFgGInaLqW7nOZvu0ZC95rtx.jpg",
			AgeRating:   13,
			TicketPrice: 53000,
		},
		{
			Title:       "Smile",
			Description: "Setelah menyaksikan kejadian aneh dan traumatis yang melibatkan seorang pasien, Dr. Rose Cotter mulai mengalami kejadian menakutkan yang tidak dapat dia jelaskan. Saat teror luar biasa mulai mengambil alih hidupnya, Rose harus menghadapi masa lalunya yang bermasalah untuk bertahan hidup dan melarikan diri dari kenyataan barunya yang mengerikan.",
			ReleaseDate: "2022-09-23",
			PosterUrl:   "https://image.tmdb.org/t/p/w500/67Myda9zANAnlS54rRjQF4dHNNG.jpg",
			AgeRating:   11,
			TicketPrice: 38000,
		},
		{
			Title:       "Encanto",
			Description: "menceritakan tentang keluarga Madrigals, sebuah keluarga yang tinggal di rumah ajaib dan masing-masing anggota keluarga memiliki keajaibannya tersendiri. Pada jaman dahulu kala, Abuela bersama suami dan anak-anaknya melarikan diri dari kerusuhan di desa.",
			ReleaseDate: "2021-10-13",
			PosterUrl:   "https://image.tmdb.org/t/p/w500/4j0PNHkMr5ax3IA8tjtxcmPU3QT.jpg",
			AgeRating:   12,
			TicketPrice: 44000,
		},
	}

	hasTable := db.Migrator().HasTable(&entities.Movies{})
	if !hasTable {
		if err := db.AutoMigrate(&entities.Movies{}); err != nil {
			return err
		}
	}

	for _, data := range listMovie {
		var movie entities.Movies
		err := db.Where(&entities.Movies{Title: data.Title}).First(&movie).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			t, err := time.Parse("2006-01-02", data.ReleaseDate)
			if err != nil {
				fmt.Println("Error parsing time:", err)
				return err
			}

			timeData := t.In(time.FixedZone("WIB", 7*60*60))

			var dataMovie = entities.Movies{
				Title:       data.Title,
				Description: data.Description,
				ReleaseDate: timeData,
				PosterUrl:   data.PosterUrl,
				AgeRating:   data.AgeRating,
				TicketPrice: data.TicketPrice,
			}

			if err := db.Create(&dataMovie).Error; err != nil {
				return err
			}
		} else {
			fmt.Println("Movie already exists:", movie.Title)
		}
	}

	return nil
}
