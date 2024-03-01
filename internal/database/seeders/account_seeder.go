package seeders

import (
	"abude-backend/internal/pkg/accounts/account"
	"abude-backend/internal/pkg/accounts/category"

	"gorm.io/gorm"
)

func AccountSeeder(db *gorm.DB) {
	if err := db.Transaction(func(tx *gorm.DB) error {
		categories := []*category.Category{
			{
				Name:        "Aset",
				Code:        "1",
				Description: "Sumber daya yang dimiliki atau dikontrol oleh perusahaan.",
				Normal:      1,
				CompanyID:   nil,
			},
			{
				Name:        "Liabilitas",
				Code:        "2",
				Description: "Kewajiban atau utang yang dimiliki oleh perusahaan kepada pihak eksternal.",
				Normal:      -1,
				CompanyID:   nil,
			},
			{
				Name:        "Pendapatan",
				Code:        "3",
				Description: "Arus masuk aset yang dihasilkan dari kegiatan rutin perusahaan.",
				Normal:      -1,
				CompanyID:   nil,
			},
			{
				Name:        "Beban",
				Code:        "4",
				Description: "Arus keluar atau pengurangan aset yang dikeluarkan dalam proses menghasilkan pendapatan.",
				Normal:      1,
				CompanyID:   nil,
			},
			{
				Name:        "Ekuitas",
				Code:        "5",
				Description: "Sisa kepentingan dalam aset perusahaan setelah dikurangi liabilitasnya.",
				Normal:      -1,
				CompanyID:   nil,
			},
		}

		if err := tx.Create(&categories).Error; err != nil {
			return err
		}

		if err := tx.Create([]*account.Account{
			{
				Name:        "Beban",
				Code:        "400000",
				Description: "Biaya beban.",
				Normal:      1,
				CategoryID:  categories[3].ID,
				CompanyID:   nil,
			},
			{
				Name:        "Beban Gaji",
				Code:        "400001",
				Description: "Biaya gaji dan upah karyawan.",
				Normal:      1,
				CategoryID:  categories[3].ID,
				CompanyID:   nil,
			},
			{
				Name:        "Beban Bahan Baku",
				Code:        "400002",
				Description: "Biaya yang harus dibayar oleh perusahaan untuk pembelian bahan baku.",
				Normal:      1,
				CategoryID:  categories[3].ID,
				CompanyID:   nil,
			},
			{
				Name:        "Beban Administrasi",
				Code:        "400003",
				Description: "Biaya-biaya administratif yang terkait dengan operasional sehari-hari perusahaan.",
				Normal:      1,
				CategoryID:  categories[3].ID,
				CompanyID:   nil,
			},
			{
				Name:        "Beban Pengiriman",
				Code:        "400004",
				Description: "Biaya yang harus dibayar oleh perusahaan untuk pengiriman barang kepada pelanggan.",
				Normal:      1,
				CategoryID:  categories[3].ID,
				CompanyID:   nil,
			},
			{
				Name:        "Beban Penyimpanan",
				Code:        "400005",
				Description: "Biaya yang terkait dengan penyimpanan barang atau inventaris perusahaan.",
				Normal:      1,
				CategoryID:  categories[3].ID,
				CompanyID:   nil,
			},
			{
				Name:        "Beban Penyusutan Perangkat Lunak",
				Code:        "400006",
				Description: "Biaya yang terkait dengan penyusutan nilai perangkat lunak yang digunakan oleh perusahaan.",
				Normal:      1,
				CategoryID:  categories[3].ID,
				CompanyID:   nil,
			},
			{
				Name:        "Beban Asuransi",
				Code:        "400007",
				Description: "Biaya-biaya yang harus dibayar oleh perusahaan untuk mendapatkan perlindungan asuransi.",
				Normal:      1,
				CategoryID:  categories[3].ID,
				CompanyID:   nil,
			},
			{
				Name:        "Beban Pajak",
				Code:        "400008",
				Description: "Biaya yang harus dibayar oleh perusahaan kepada pemerintah sebagai pajak.",
				Normal:      1,
				CategoryID:  categories[3].ID,
				CompanyID:   nil,
			},
			{
				Name:        "Beban Penanganan Produk Kembali",
				Code:        "400009",
				Description: "Biaya yang terkait dengan penanganan atau pengembalian produk yang dibeli oleh pelanggan.",
				Normal:      1,
				CategoryID:  categories[3].ID,
				CompanyID:   nil,
			},
			{
				Name:        "Beban Biaya Penjualan",
				Code:        "400010",
				Description: "Biaya yang terkait dengan penjualan produk atau jasa perusahaan.",
				Normal:      1,
				CategoryID:  categories[3].ID,
				CompanyID:   nil,
			},
		}).Error; err != nil {
			return err
		}

		if err := tx.Create([]*account.Account{
			{
				Name:        "Pendapatan Penjualan",
				Code:        "300001",
				Description: "Pendapatan yang dihasilkan oleh perusahaan dari penjualan barang atau jasa kepada pelanggan.",
				Normal:      -1,
				CategoryID:  categories[2].ID,
				CompanyID:   nil,
			},
			{
				Name:        "Pendapatan Sewa",
				Code:        "300002",
				Description: "Pendapatan yang dihasilkan oleh perusahaan dari penyewaan properti atau aset lainnya.",
				Normal:      -1,
				CategoryID:  categories[2].ID,
				CompanyID:   nil,
			},
			{
				Name:        "Pendapatan Bunga",
				Code:        "300003",
				Description: "Pendapatan yang dihasilkan oleh perusahaan dari bunga yang diterima dari investasi atau pinjaman.",
				Normal:      -1,
				CategoryID:  categories[2].ID,
				CompanyID:   nil,
			},
			{
				Name:        "Pendapatan Royalti",
				Code:        "300004",
				Description: "Pendapatan yang dihasilkan oleh perusahaan dari royalti yang diterima atas penggunaan hak cipta, paten, atau merek dagang.",
				Normal:      -1,
				CategoryID:  categories[2].ID,
				CompanyID:   nil,
			},
			{
				Name:        "Pendapatan Bunga Deposito",
				Code:        "300005",
				Description: "Pendapatan yang dihasilkan oleh perusahaan dari bunga yang diterima dari simpanan deposito di bank atau lembaga keuangan lainnya.",
				Normal:      -1,
				CategoryID:  categories[2].ID,
				CompanyID:   nil,
			},
			{
				Name:        "Pendapatan Dividen",
				Code:        "300006",
				Description: "Pendapatan yang dihasilkan oleh perusahaan dari dividen yang diterima dari investasi di perusahaan lain.",
				Normal:      -1,
				CategoryID:  categories[2].ID,
				CompanyID:   nil,
			},
			{
				Name:        "Pendapatan Komisi",
				Code:        "300007",
				Description: "Pendapatan yang dihasilkan oleh perusahaan dari komisi yang diterima dari penjualan produk atau jasa lainnya.",
				Normal:      -1,
				CategoryID:  categories[2].ID,
				CompanyID:   nil,
			},
			{
				Name:        "Pendapatan Lisensi",
				Code:        "300008",
				Description: "Pendapatan yang dihasilkan oleh perusahaan dari biaya lisensi yang diterima atas penggunaan hak cipta, paten, atau merek dagang.",
				Normal:      -1,
				CategoryID:  categories[2].ID,
				CompanyID:   nil,
			},
			{
				Name:        "Pendapatan Layanan",
				Code:        "300009",
				Description: "Pendapatan yang dihasilkan oleh perusahaan dari penyediaan layanan kepada pelanggan.",
				Normal:      -1,
				CategoryID:  categories[2].ID,
				CompanyID:   nil,
			},
			{
				Name:        "Pendapatan Lainnya",
				Code:        "300010",
				Description: "Pendapatan lainnya yang dihasilkan oleh perusahaan dan tidak termasuk dalam kategori pendapatan lainnya.",
				Normal:      -1,
				CategoryID:  categories[2].ID,
				CompanyID:   nil,
			},
		}).Error; err != nil {
			return err
		}

		if err := tx.Create([]*account.Account{
			{
				Name:        "Kas",
				Code:        "100001",
				Description: "Aset yang dimiliki oleh perusahaan dalam bentuk uang tunai atau investasi yang mudah dicairkan menjadi uang tunai dengan cepat.",
				Normal:      1,
				CategoryID:  categories[0].ID,
				CompanyID:   nil,
			},
			{
				Name:        "Piutang Usaha",
				Code:        "100002",
				Description: "Hak klaim atas pembayaran yang harus diterima oleh perusahaan dari pelanggan atas penjualan barang atau jasa.",
				Normal:      1,
				CategoryID:  categories[0].ID,
				CompanyID:   nil,
			},
			{
				Name:        "Persediaan Barang",
				Code:        "100003",
				Description: "Barang-barang yang dimiliki oleh perusahaan untuk dijual dalam operasi bisnisnya.",
				Normal:      1,
				CategoryID:  categories[0].ID,
				CompanyID:   nil,
			},
			{
				Name:        "Aset Tetap",
				Code:        "100004",
				Description: "Aset jangka panjang yang dimiliki oleh perusahaan dan digunakan dalam operasi bisnisnya, seperti tanah, bangunan, mesin, dan peralatan.",
				Normal:      1,
				CategoryID:  categories[0].ID,
				CompanyID:   nil,
			},
			{
				Name:        "Investasi Jangka Panjang",
				Code:        "100005",
				Description: "Investasi dalam bentuk surat berharga atau kepemilikan saham di perusahaan lain yang dimiliki untuk jangka waktu yang lebih dari satu tahun.",
				Normal:      1,
				CategoryID:  categories[0].ID,
				CompanyID:   nil,
			},
			{
				Name:        "Aset Lainnya",
				Code:        "100006",
				Description: "Aset lainnya yang dimiliki oleh perusahaan dan tidak termasuk dalam kategori aset lainnya.",
				Normal:      1,
				CategoryID:  categories[0].ID,
				CompanyID:   nil,
			},
		}).Error; err != nil {
			return err
		}

		if err := tx.Create([]*account.Account{
			{
				Name:        "Liabilitas",
				Code:        "200000",
				Description: "Jumlah uang yang dipinjam oleh perusahaan.",
				Normal:      -1,
				CategoryID:  categories[1].ID,
				CompanyID:   nil,
			},
			{
				Name:        "Utang Bank",
				Code:        "200001",
				Description: "Jumlah uang yang dipinjam oleh perusahaan dari bank atau lembaga keuangan lainnya.",
				Normal:      -1,
				CategoryID:  categories[1].ID,
				CompanyID:   nil,
			},
			{
				Name:        "Utang Usaha",
				Code:        "200002",
				Description: "Utang yang harus dibayar oleh perusahaan kepada pihak lain sebagai bagian dari kegiatan bisnisnya.",
				Normal:      -1,
				CategoryID:  categories[1].ID,
				CompanyID:   nil,
			},
			{
				Name:        "Utang Gaji dan Upah",
				Code:        "200003",
				Description: "Jumlah gaji dan upah yang harus dibayar oleh perusahaan kepada karyawan.",
				Normal:      -1,
				CategoryID:  categories[1].ID,
				CompanyID:   nil,
			},
			{
				Name:        "Utang Pajak",
				Code:        "200004",
				Description: "Utang yang harus dibayar oleh perusahaan kepada pemerintah sebagai pajak.",
				Normal:      -1,
				CategoryID:  categories[1].ID,
				CompanyID:   nil,
			},
			{
				Name:        "Utang Jangka Panjang",
				Code:        "200005",
				Description: "Utang yang harus dibayar oleh perusahaan dalam jangka waktu lebih dari satu tahun.",
				Normal:      -1,
				CategoryID:  categories[1].ID,
				CompanyID:   nil,
			},
			{
				Name:        "Liabilitas Lainnya",
				Code:        "200006",
				Description: "Liabilitas lainnya yang dimiliki oleh perusahaan dan tidak termasuk dalam kategori liabilitas lainnya.",
				Normal:      -1,
				CategoryID:  categories[1].ID,
				CompanyID:   nil,
			},
		}).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		panic(err)
	}
}
