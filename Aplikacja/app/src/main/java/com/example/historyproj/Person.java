package com.example.historyproj;

import com.google.gson.annotations.SerializedName;
import java.util.List;

public class Person {
    @SerializedName("id")
    private int id;

    @SerializedName("name")
    private String name;

    @SerializedName("inscription")
    private String inscription;

    @SerializedName("other_names")
    private List<String> otherNames;

    @SerializedName("code_names")
    private List<String> codeNames;

    @SerializedName("birth_date")
    private String birthDate;

    @SerializedName("birth_place")
    private String birthPlace;

    @SerializedName("death_date")
    private String deathDate;

    @SerializedName("death_place")
    private String deathPlace;

    @SerializedName("burial_place")
    private String burialPlace;

    @SerializedName("cemetery")
    private String cemetery;

    @SerializedName("quarter")
    private String quarter;

    @SerializedName("row")
    private String row;

    @SerializedName("grave")
    private String grave;

    @SerializedName("ranks")
    private String ranks;

    @SerializedName("badges")
    private String badges;

    @SerializedName("activity")
    private String activity;

    @SerializedName("description")
    private String description;

    @SerializedName("sources")
    private String sources;

    // Getters and setters for all fields

    public int getId() {
        return id;
    }

    public String getName() {
        return name;
    }

    public String getInscription() {
        return inscription;
    }

    public List<String> getOtherNames() {
        return otherNames;
    }

    public List<String> getCodeNames() {
        return codeNames;
    }

    public String getBirthDate() {
        return birthDate;
    }

    public String getBirthPlace() {
        return birthPlace;
    }

    public String getDeathDate() {
        return deathDate;
    }

    public String getDeathPlace() {
        return deathPlace;
    }

    public String getBurialPlace() {
        return burialPlace;
    }

    public String getCemetery() {
        return cemetery;
    }

    public String getQuarter() {
        return quarter;
    }

    public String getRow() {
        return row;
    }

    public String getGrave() {
        return grave;
    }

    public String getRanks() {
        return ranks;
    }

    public String getBadges() {
        return badges;
    }

    public String getActivity() {
        return activity;
    }

    public String getDescription() {
        return description;
    }

    public String getSources() {
        return sources;
    }
}